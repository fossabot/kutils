package pkg

import (
	"fmt"
	"k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sort"
	"strings"
)

type Ingress struct {
	Name string
	Hosts []string
	Class string
	ServiceName string
	ServicePort int32
	PathTypePrefix v1beta1.PathType
	RewriteWWW bool
}

func (ingress *Ingress) GetHosts() (hosts []string) {
	if ingress.RewriteWWW {
		hosts = ingress.HostsWWW()
	} else {
		hosts = ingress.Hosts
	}
	sort.Strings(hosts)
	return
}

func (ingress *Ingress) NginxRewriteSnippet() string {
	var rewriteHosts []string
	for _, host := range ingress.Hosts {
		if !strings.HasPrefix(host, "www.") {
			rewriteHosts = append(rewriteHosts, host)
		}
	}
	return NewNginxWWWRewriteRules(rewriteHosts...)
}

func (ingress *Ingress) HostsWWW() (hosts []string){
	for _, host := range ingress.Hosts {
		if !strings.HasPrefix(host, "www.") {
			www := fmt.Sprintf("www.%v", host)
			if !StringInSlice(www, ingress.Hosts) {
				hosts = append(hosts, www)
			}
		}
		hosts = append(hosts, host)
	}
	return
}

func (ingress *Ingress) TLS() (tls []v1beta1.IngressTLS) {
	for _, host := range ingress.GetHosts() {
		tls = append(tls, v1beta1.IngressTLS{
			Hosts: 		[]string{host},
			SecretName: fmt.Sprintf("tls-%v", SlugifyString(host)),
		})
	}
	return
}

func (ingress *Ingress) Rules() (rules []v1beta1.IngressRule) {
	for _, host := range ingress.GetHosts() {
		rules = append(rules, v1beta1.IngressRule{
			Host:             host,
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: &v1beta1.HTTPIngressRuleValue{
					Paths: []v1beta1.HTTPIngressPath{
						{
							Path:     "/",
							PathType: &ingress.PathTypePrefix,
							Backend:  v1beta1.IngressBackend{
								ServiceName: ingress.ServiceName,
								ServicePort: intstr.IntOrString{
									IntVal: ingress.ServicePort,
									StrVal: string(ingress.ServicePort),
								},
							},
						},
					},
				},
			},
		})
	}
	return
}

func (ingress *Ingress) KubernetesObject() *v1beta1.Ingress {
	var ingressTls []v1beta1.IngressTLS
	for _, host := range ingress.Hosts {
		ingressTls = append(ingressTls, v1beta1.IngressTLS{
			Hosts: []string{host},
			SecretName: fmt.Sprintf("tls-%v", SlugifyString(host)),
		})
	}
	object := &v1beta1.Ingress{
		TypeMeta: v1.TypeMeta{
			APIVersion: "networking.k8s.io/v1beta1",
			Kind:       "Ingress",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: ingress.Name,
			Annotations: map[string]string{},
		},
		Spec: v1beta1.IngressSpec{
			IngressClassName: &ingress.Class,
			Backend: &v1beta1.IngressBackend{
				ServiceName: ingress.ServiceName,
				ServicePort: intstr.IntOrString{
					StrVal: string(ingress.ServicePort),
					IntVal: ingress.ServicePort,
				},
			},
			TLS: ingress.TLS(),
			Rules: ingress.Rules(),
		},
		Status: v1beta1.IngressStatus{},
	}
	if ingress.RewriteWWW {
		object.ObjectMeta.Annotations["nginx.ingress.kubernetes.io/configuration-snippet"] = ingress.NginxRewriteSnippet()
	}
	return object
}
