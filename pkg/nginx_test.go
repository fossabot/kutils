package pkg

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
	"testing"
)

func TestNewNginxWWWRewriteRule(t *testing.T) {
	host := "example.com"
	rule := NewNginxWWWRewriteRule(host)
	AssertStringEqual(t,
		rule,
		"if ($host = 'www.example.com') {\nrewrite ^ https://example.com$request_uri permanent;\n}")
}

func TestNewNginxWWWRewriteRules(t *testing.T) {
	rule := NewNginxWWWRewriteRules("example.com", "example2.com", "example3.com")
	AssertStringEqual(t,
		rule,
		"if ($host = 'www.example.com') {\nrewrite ^ https://example.com$request_uri permanent;\n}\n" +
		"if ($host = 'www.example2.com') {\nrewrite ^ https://example2.com$request_uri permanent;\n}\n" +
		"if ($host = 'www.example3.com') {\nrewrite ^ https://example3.com$request_uri permanent;\n}")
}

// todo live test with controller
func TestNginxWWWRewriteRuleConfigurationSnippet(t *testing.T) {
	expected := "annotations:\n" +
		"  nginx.ingress.kubernetes.io/configuration-snippet: |-\n" +
		"    if ($host = 'www.s1.com') {\n" +
		"    rewrite ^ https://s1.com$request_uri permanent;\n" +
		"    }\n" +
		"    if ($host = 'www.s2.pl') {\n" +
		"    rewrite ^ https://s2.pl$request_uri permanent;\n" +
		"    }\n" +
		"    if ($host = 'www.s3.it') {\n" +
		"    rewrite ^ https://s3.it$request_uri permanent;\n" +
		"    }\n" +
		"creationTimestamp: null\n" +
		"name: test\n"
	snippet := NewNginxWWWRewriteRules("s1.com", "s2.pl", "s3.it")
	meta := v1.ObjectMeta{
		Name:                       "test",
		Annotations: map[string]string{
			"nginx.ingress.kubernetes.io/configuration-snippet": snippet,
		},
	}
	if data, err := yaml.Marshal(meta); err != nil {
		t.Error(err)
	} else {
		AssertStringEqual(t, string(data), expected)
	}
}