package goformer

import "reflect"

// this writes a `value` byte to every property
// I recommend using >0 values
// DO NOT use any properties from the returned struct
func NewUnsafeDummy[T any](value byte) T {
	var t T
	valueOf := reflect.ValueOf(&t).Elem()

	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		addr := field.Addr().UnsafePointer()
		*(*byte)(addr) = value
	}

	return t
}

func compareEqual(a reflect.Value, b reflect.Value) bool {
	addrA := a.Addr().UnsafePointer()
	addrB := b.Addr().UnsafePointer()

	return *(*byte)(addrA) == *(*byte)(addrB)
}

// compares 2 structs by looping over `a` properties
// and comparing it to the same name `b` property
func CompareDummiesEqual(a any, b any) bool {

	valueOfA := reflect.ValueOf(&a).Elem()
	valueOfB := reflect.ValueOf(&b).Elem()

	for i := 0; i < valueOfA.NumField(); i++ {
		fieldA := valueOfA.Field(i)
		fieldB := valueOfB.FieldByName(valueOfA.Type().Field(i).Name)

		if !compareEqual(fieldA, fieldB) {
			return false
		}
	}

	return true
}
