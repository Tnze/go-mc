package nbt

import (
	"bytes"
	"fmt"
)

//goland:noinspection SpellCheckingInspection
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

	data, err := Marshal(value)
	if err != nil {
		panic(err)
	}

	fmt.Printf("% 02x ", data)

	// Output:
	//	0a 00 00 08 00 04 6e 61 6d 65 00 04 54 6e 7a 65 00
}

func ExampleEncoder_WriteSNBT() {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).WriteSNBT(`{ name: [Tnze, "Xi_Xi_Mi"]}`); err != nil {
		panic(err)
	}
	fmt.Printf("% 02x ", buf.Bytes())

	// Output:
	// 0a 00 00 09 00 04 6e 61 6d 65 08 00 00 00 02 00 04 54 6e 7a 65 00 08 58 69 5f 58 69 5f 4d 69 00
}
