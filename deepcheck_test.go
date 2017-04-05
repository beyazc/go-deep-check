package deepcheck

import (
	"testing"
)

type student struct {
	name string
	age  int
}

func TestDifferentTypes(t *testing.T) {
	student1 := student{}
	student2 := student{}
	number := 12
	if Equals(number, student2) {
		t.Error("different types should not be equal")
	}

	if !Equals(student1, student2) {
		t.Error("same type should be equal")
	}

	if Equals(student1, &student1) {
		t.Error("a type with its reference is also not same")
	}

}

func TestNil(t *testing.T) {
	student1 := student{}
	number := 12
	if Equals(number, nil) {
		t.Error("only one of the parameter can not be nil")
	}
	if Equals(nil, student1) {
		t.Error("only one of the parameter can not be nil")
	}
	if !Equals(nil, nil) {
		t.Error("two nils are equal")
	}

}

func TestEqualsInterfacePointer(t *testing.T) {

}
