package main

import (
	"github.com/shamaton/msgpack"
)

func main() {
	type Struct struct {
		String string
	}
	v := Struct{String: "msgpack"}

	d, err := msgpack.Encode(v)
	if err != nil {
		panic(err)
	}
	r := Struct{}
	err = msgpack.Decode(d, &r)
	if err != nil {
		panic(err)
	}
}
