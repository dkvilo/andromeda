package functions

import "reflect"

// SizeOf any type
func SizeOf(t interface{}) uintptr {
	return reflect.TypeOf(t).Size()
}
