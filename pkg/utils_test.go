package pkg

import (
	"os"
	"testing"
)

func TestSlugifyString(t *testing.T) {
	AssertStringEqual(t, SlugifyString("www.example.com"), "www-example-com")
}

func TestSaveYaml(t *testing.T) {
	filePath := "./__testfile.yaml"
	type name struct {
		A string
		B int
		C map[string]string
	}
	obj := &name{
		A: "1",
		B: 2,
		C: map[string]string{"a": "b"},
	}
	SaveYaml(obj, filePath)
	err := os.Remove(filePath)
	if err != nil {t.Error(err)}
}

func TestReadYaml(t *testing.T) {
	filePath := "./__testfile.yaml"
	type name struct {
		A string
		B int
	}
	obj := &name{
		A: "1",
		B: 2,
	}
	SaveYaml(obj, filePath)

	var obj2 name
	ReadYaml(filePath, &obj2)
	AssertStringEqual(t, obj.A, obj2.A)
	AssertIntEqual(t, obj.B, obj2.B)
	err := os.Remove(filePath)
	if err != nil {t.Error(err)}
}

func TestStringInSlice(t *testing.T) {
	slice := []string{"test", "test2"}
	str := "test"
	if !StringInSlice(str, slice) {
		t.Error("String not in slice!")
	}
	if StringInSlice("test3", slice) {
		t.Error("String in slice!")
	}
}
//
//func TestValidateHost(t *testing.T) {
//	if ValidateHost("3") != false {
//		t.Error("Host validation failed")
//	}
//	if ValidateHost("example.com") == false {
//		t.Error("Host validation failed")
//	}
//}