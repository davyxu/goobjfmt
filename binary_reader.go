package goobjfmt

import (
	"encoding/binary"
	"reflect"
)

func BinaryRead(data []byte, obj interface{}) error {

	v := reflect.ValueOf(obj)

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	}

	d := &decoder{order: binary.LittleEndian, buf: data}
	d.value(v)

	return nil
}
