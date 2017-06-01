package goobjfmt

import (
	"reflect"
)

func dataSize(v reflect.Value) int {

	switch v.Kind() {
	case reflect.Array:
		if s := dataSize(v.Elem()); s >= 0 {
			return s*v.Type().Len() + 4
		}
	case reflect.Slice:
		l := v.Len()
		elemSize := int(v.Type().Elem().Size())
		return l*elemSize + 4

	case reflect.String:
		t := v.Len()
		return t + 4
	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64:
		return int(v.Type().Size())
	case reflect.Struct:
		sum := 0
		for i := 0; i < v.NumField(); i++ {

			s := dataSize(v.Field(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum
	}

	return -1
}

func BinarySize(obj interface{}) int {
	v := reflect.Indirect(reflect.ValueOf(obj))
	return dataSize(v)
}
