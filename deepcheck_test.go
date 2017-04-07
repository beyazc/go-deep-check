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
	intVal1 := 2
	intVal2 := 2
	intval3 := 3

	if !Equals(intVal1, intVal2) {
		t.Error("same int values shoul be equal")
	}

	if Equals(intVal1, intval3) {
		t.Error("different int values shoul not be equal")
	}

	floatVal1 := 2.0
	floatVal2 := 2.0
	floatVal3 := 3

	if !Equals(floatVal1, floatVal2) {
		t.Error("same float values shoul be equal")
	}

	if Equals(floatVal1, floatVal3) {
		t.Error("different float values shoul not be equal")
	}

	boolVal1 := true
	boolVal2 := true
	boolVal3 := false

	if !Equals(boolVal1, boolVal2) {
		t.Error("same bool values shoul be equal")
	}

	if Equals(boolVal1, boolVal3) {
		t.Error("different bool values shoul not be equal")
	}

	stringVal1 := "deepcheck"
	stringVal2 := "deepcheck"
	stringVal3 := "checdeep"

	if !Equals(stringVal1, stringVal2) {
		t.Error("same bool string shoul be equal")
	}

	if Equals(stringVal1, stringVal3) {
		t.Error("different string values shoul not be equal")
	}

	uintVal1 := 1
	uintVal2 := 1
	uintVal3 := 2
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
