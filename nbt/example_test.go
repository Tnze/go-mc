package nbt

import (
	"bytes"
	"fmt"
)

func ExampleDecoder_Decode() {
	reader := bytes.NewReader([]byte{
		0x0a, // Start TagCompound("")
		0x00, 0x00,

		0x08, // TagString("Author"): "Tnze"
		0x00, 0x06, 'A', 'u', 't', 'h', 'o', 'r',
		0x00, 0x04, 'T', 'n', 'z', 'e',

		0x00, // End TagCompound
	})

	var value struct {
		Author string
	}

	decoder := NewDecoder(reader)
	_, err := decoder.Decode(&value)
	if err != nil {
		panic(err)
	}

	fmt.Println(value)

	// Output:
	// {Tnze}
}

func ExampleDecoder_Decode_singleTagString() {
	reader := bytes.NewReader([]byte{
		// TagString
		0x08,
		// TagName
		0x00, 0x04,
		0x6e, 0x61, 0x6d, 0x65,
		// Content
		0x00, 0x09,
		0x42, 0x61, 0x6e, 0x61, 0x6e, 0x72, 0x61, 0x6d, 0x61,
	})

	var Name string

	decoder := NewDecoder(reader)
	tagName, err := decoder.Decode(&Name)
	if err != nil {
		panic(err)
	}

	fmt.Println(tagName)
	fmt.Println(Name)

	// Output:
	// name
	// Bananrama
}

func ExampleEncoder_Encode_tagCompound() {
	value := struct {
		Name string `nbt:"name"`
	}{"Tnze"}

	var buff bytes.Buffer
	encoder := NewEncoder(&buff)
	err := encoder.Encode(value, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("% 02x ", buff.Bytes())

	// Output:
	//	0a 00 00 08 00 04 6e 61 6d 65 00 04 54 6e 7a 65 00
}

func ExampleEncoder_writeSNBT() {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(StringifiedMessage(`{ name: [Tnze, "Xi_Xi_Mi"]}`), ""); err != nil {
		panic(err)
	}
	fmt.Printf("% 02x ", buf.Bytes())

	// Output:
	// 0a 00 00 09 00 04 6e 61 6d 65 08 00 00 00 02 00 04 54 6e 7a 65 00 08 58 69 5f 58 69 5f 4d 69 00
}
