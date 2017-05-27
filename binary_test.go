package goobjfmt

import (
	"fmt"
	"testing"
)

type P struct {
	Name    string
	X, Y, Z int
}

func TestWrite(t *testing.T) {

	data, err := BinaryWrite(P{"hello", 3, 4, 5})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	var p2 P
	err = BinaryRead(data, &p2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p2)
	}
}
