package pkg

import (
	"io/ioutil"
	"sigs.k8s.io/yaml"
	"strings"
)

func SlugifyString(s string) string {
	s = strings.Replace(s, ".", "-", -1)
	return s
}

func SaveYaml(o interface{}, filePath string)  (err error) {
	yamlBytes, err := yaml.Marshal(o)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, yamlBytes, 0644)
	if err != nil {
		return err
	}
	return
}

func ReadYaml(filePath string, o interface{}) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil{ panic(err)}
	err = yaml.Unmarshal(file, o)
	if err != nil {panic(err)}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//
//func ValidateHost(s string) bool {
//	_, err := url.ParseRequestURI(s)
//	if err != nil {
//		fmt.Print(err)
//		return false
//	}
//	return true
//}
