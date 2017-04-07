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

	//two nils are equal
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		log.Println(ErrNilParameterReceived)
		return false
	}
	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)
	return equalsInternal(aVal, bVal)
}

func equalsInternal(aVal, bVal reflect.Value) bool {
	aType := aVal.Type()
	bType := bVal.Type()

	if aType != bType {
		log.Println(ErrTypesDoesNotMatch)
		return false
	}

	if !aVal.IsValid() && !bVal.IsValid() {
		return true
	}

	if !aVal.IsValid() && bVal.IsValid() {
		log.Println(ErrNilParameterReceived)
		return false
	}

	if !bVal.IsValid() && aVal.IsValid() {
		log.Println(ErrNilParameterReceived)
		return false
	}

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
		aInt := aVal.Int()
		bInt := bVal.Int()
		if aInt != bInt {
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

	//struct types
	case reflect.Struct:

		for i := 0; i < aType.NumField(); i++ {
			fieldA := (aVal.Field(i))
			fieldB := (bVal.Field(i))

			if fieldComparisionRes := equalsInternal(fieldA, fieldB); !fieldComparisionRes {
				fmt.Printf("fields did not matched for field %v,%v\n", fieldA, fieldB)
				return false
			}
		}
		return true

	default:
		fmt.Println(ErrNotHandled)
		return false
	}

	return true
}

func fieldEquals(a, b interface{}) bool {

	return false
}
