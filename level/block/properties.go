package block

import (
	"strconv"
)

type (
	Boolean   bool
	Direction string
	Integer   int
)

func (b Boolean) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(bool(b))), nil
}

func (b *Boolean) UnmarshalText(text []byte) (err error) {
	*((*bool)(b)), err = strconv.ParseBool(string(text))
	return
}

func (d Direction) MarshalText() (text []byte, err error) {
	return []byte(d), nil
}

func (d *Direction) UnmarshalText(text []byte) error {
	*d = Direction(text)
	return nil
}

func (i Integer) MarshalText() (text []byte, err error) {
	return []byte(strconv.Itoa(int(i))), nil
}

func (i *Integer) UnmarshalText(text []byte) (err error) {
	*((*int)(i)), err = strconv.Atoi(string(text))
	return
}
