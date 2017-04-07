package deepcheck

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var (
	//ErrTypesDoesNotMatch when types are different
	ErrTypesDoesNotMatch = errors.New("variables type did not matched")
	//ErrNilParameterReceived indicates that a compare parameter was nil
	ErrNilParameterReceived = errors.New("A nil parameter recevied")
	//ErrValueMismatched values does not match
	ErrValueMismatched = errors.New("Values does not match")
	//ErrNotHandled indicates that this comparision is not handled yet
	ErrNotHandled = errors.New("comparision for type is not defined")
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
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)

	switch aVal.Kind() {
	//primitive types
	case reflect.Float32, reflect.Float64:
		if aVal.Float() != bVal.Float() {
			log.Println(ErrValueMismatched)
			return false
		}
	case reflect.Bool:
		if aVal.Bool() != bVal.Bool() {
			log.Println(ErrValueMismatched)
			return false
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if aVal.Int() != bVal.Int() {
			log.Println(ErrValueMismatched)
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if aVal.Uint() != bVal.Uint() {
			log.Println(ErrValueMismatched)
			return false
		}
	case reflect.String:
		if aVal.String() != bVal.String() {
			log.Println(ErrValueMismatched)
			return false
		}

	default:
		fmt.Println(ErrNotHandled)
		return false
	}

	return true
}

func fieldEquals(a, b interface{}) bool {

	return false
}
