package deepcheck

import (
	"errors"
	"log"
	"reflect"
)

var (
	//ErrTypesDoesNotMatch when types are different
	ErrTypesDoesNotMatch = errors.New("variables type did not matched")
	//ErrNilParameterReceived indicates that a compare parameter was nil
	ErrNilParameterReceived = errors.New("A nil parameter recevied")
)

//Equals returns whether given interfaces are equal or not
func Equals(a, b interface{}) bool {
	//aVal := reflect.Value(a)
	//bVal := reflect.Value(b)

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	if aType != bType {
		log.Println(ErrTypesDoesNotMatch)
		return false
	}

	if a == nil && b == nil {
		return true
	}

	if a == nil && b != nil {
		log.Println(ErrNilParameterReceived)
		return false
	}

	if b == nil && a != nil {
		log.Println(ErrNilParameterReceived)
		return false
	}

	return true
}

func fieldEquals(a, b interface{}) bool {

	return false
}
