package nbt

import (
	"bytes"
	"fmt"
)

func ExampleUnmarshal() {
	var data = []byte{
		0x08, 0x00, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x00, 0x09,
		0x42, 0x61, 0x6e, 0x61, 0x6e, 0x72, 0x61, 0x6d, 0x61,
	}

	var Name string

	if err := Unmarshal(data, &Name); err != nil {
		panic(err)
	}

	fmt.Println(Name)

	// Output: Bananrama
}

func ExampleMarshal() {
	var value = struct {
		Name string `nbt:"name"`
	}{"Tnze"}

	var buf bytes.Buffer
	if err := Marshal(&buf, value); err != nil {
		panic(err)
	}

	fmt.Printf("% 02x ", buf.Bytes())

	// Output:
	//	0a 00 00 08 00 00 00 04 54 6e 7a 65 00
}
