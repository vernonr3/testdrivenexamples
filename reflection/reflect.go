package main

import (
	"reflect"
)

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	// kind returns the kind of value val is. as an uint8 value so it can be easily compared to a list of constants..
	if val.Kind() == reflect.Ptr {
		//Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Pointer. It returns the zero Value if v is nil.
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, fn func(input string)) {
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			// index returns vals i'th element
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			// Field returns the field to which the value refers. It panics if val is not a struct or i is out of range
			walkValue(val.Field(i))
		}
	case reflect.Map:
		// interface() returns the value as an interface
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.String:
		// string() returns the value as an string
		fn(val.String())
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		// Call calls the function - in this case with no arguments.
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}
