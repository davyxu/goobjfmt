package goobjfmt

import (
	"encoding/binary"
	"errors"
	"reflect"
)

func BinaryWrite(obj interface{}) ([]byte, error) {

	// Fallback to reflect-based encoding.
	v := reflect.Indirect(reflect.ValueOf(obj))
	size := dataSize(v)
	if size < 0 {
		return nil, errors.New("invalid type " + reflect.TypeOf(obj).String())
	}

	buf := make([]byte, size)

	e := &encoder{order: binary.LittleEndian, buf: buf}
	e.value(v)

	return buf, nil
}
