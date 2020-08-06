package pkg

import (
	"testing"
)

func AssertStringEqual(t *testing.T, s string, e string)  {
	if s != e {
		t.Errorf("%v not equal to %v", s, e)
	}
}

func AssertIntEqual(t *testing.T, s int, e int)  {
	if s != e {
		t.Errorf("%v not equal to %v", s, e)
	}
}


func TestAssertStringEqual(t *testing.T)  {
	AssertStringEqual(t, "2", "2")
}


func TestAssertIntEqual(t *testing.T)  {
	AssertIntEqual(t, 2, 2)
}

