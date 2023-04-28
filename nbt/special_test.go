package nbt_test

import (
	"fmt"
	"testing"

	"github.com/Tnze/go-mc/nbt"
)

func ExampleMarshal_anonymousStructField() {
	type A struct{ F string }
	type B struct{ E string }
	type S struct {
		A           // anonymous fields are usually marshaled as if their inner exported fields were fields in the outer struct
		B `nbt:"B"` // anonymous field, but with an explicit tag name specified
	}

	var val S
	val.F = "Tnze"
	val.E = "GoMC"

	data, err := nbt.Marshal(val)
	if err != nil {
		panic(err)
	}

	var snbt nbt.StringifiedMessage
	if err := nbt.Unmarshal(data, &snbt); err != nil {
		panic(err)
	}
	fmt.Println(snbt)

	// Output:
	// {F:Tnze,B:{E:GoMC}}
}

func ExampleUnmarshal_anonymousStructField() {
	type A struct{ F string }
	type B struct{ E string }
	type S struct {
		A           // anonymous fields are usually marshaled as if their inner exported fields were fields in the outer struct
		B `nbt:"B"` // anonymous field, but with an explicit tag name specified
	}

	data, err := nbt.Marshal(nbt.StringifiedMessage(`{F:Tnze,B:{E:GoMC}}`))
	if err != nil {
		panic(err)
	}

	var val S
	if err := nbt.Unmarshal(data, &val); err != nil {
		panic(err)
	}
	fmt.Println(val.F)
	fmt.Println(val.E)

	// Output:
	// Tnze
	// GoMC
}

func TestMarshal_anonymousPointerNesting(t *testing.T) {
	type A struct{ T string }
	type B struct{ *A }
	type C struct{ B }

	val := C{B{&A{"Tnze"}}}

	data, err := nbt.Marshal(val)
	if err != nil {
		panic(err)
	}

	var snbt nbt.StringifiedMessage
	if err := nbt.Unmarshal(data, &snbt); err != nil {
		panic(err)
	}
	want := `{T:Tnze}`
	if string(snbt) != want {
		t.Errorf("Marshal nesting anonymous struct error, got %q, want %q", snbt, want)
	}
}

func TestMarshal_anonymousNonStruct(t *testing.T) {
	type A [3]int32
	type B struct{ *A }
	type C struct{ B }

	val := C{B{&A{0, -1, 3}}}

	data, err := nbt.Marshal(val)
	if err != nil {
		panic(err)
	}

	var snbt nbt.StringifiedMessage
	if err := nbt.Unmarshal(data, &snbt); err != nil {
		panic(err)
	}
	want := `{A:[I;0I,-1I,3I]}`
	if string(snbt) != want {
		t.Errorf("Marshal nesting anonymous struct error, got %q, want %q", snbt, want)
	}
}
