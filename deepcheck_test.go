package deepcheck

import (
	"testing"
)

type student struct {
	name string
	age  int
}

func TestDifferentTypes(t *testing.T) {
	student1 := student{"john", 10}
	student2 := student{"john", 10}
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

	if !Equals(student1, student1) {
		t.Error("same struct must be equal its equivelent")
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
	number1 := 12
	number2 := 13

	if Equals(number1, number2) {
		t.Error("different values should not be equal")
	}

	number2 = 12
	if !Equals(number1, number2) {
		t.Error("same values should be equal")
	}

}

func TestPrimitives(t *testing.T) {
	var intVal1 int = 2
	var intVal2 int = 2
	var intval3 int = 3

	if !Equals(intVal1, intVal2) {
		t.Error("same int values shoul be equal")
	}

	if Equals(intVal1, intval3) {
		t.Error("different int values shoul not be equal")
	}

	var floatVal1 float32 = 2.0
	var floatVal2 float32 = 2.0
	var floatVal3 float32 = 3

	if !Equals(floatVal1, floatVal2) {
		t.Error("same float values shoul be equal")
	}

	if Equals(floatVal1, floatVal3) {
		t.Error("different float values shoul not be equal")
	}

	var boolVal1 bool = true
	var boolVal2 bool = true
	var boolVal3 bool = false

	if !Equals(boolVal1, boolVal2) {
		t.Error("same bool values shoul be equal")
	}

	if Equals(boolVal1, boolVal3) {
		t.Error("different bool values shoul not be equal")
	}

	var stringVal1 string = "deepcheck"
	var stringVal2 string = "deepcheck"
	var stringVal3 string = "checdeep"

	if !Equals(stringVal1, stringVal2) {
		t.Error("same bool string shoul be equal")
	}

	if Equals(stringVal1, stringVal3) {
		t.Error("different string values shoul not be equal")
	}

	var uintVal1 uint = 1
	var uintVal2 uint = 1
	var uintVal3 uint = 2
	if !Equals(uintVal1, uintVal2) {
		t.Error("same uint shoul be equal")
	}

	if Equals(uintVal1, uintVal3) {
		t.Error("different uint values shoul not be equal")
	}

}

func TestStruc(t *testing.T) {
	mystudent1 := student{"john", 10}
	mystudent2 := student{"john", 10}
	mystudent3 := student{"mark", 20}

	if !Equals(mystudent1, mystudent2) {
		t.Error("same struct shoul be equal")
	}

	if Equals(mystudent1, mystudent3) {
		t.Error("different struct values shoul not be equal")
	}

}
