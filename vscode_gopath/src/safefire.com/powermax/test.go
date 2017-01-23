package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/ugorji/go/codec"
)

type value struct {
	Tm     uint64
	Value  int
	Status int
}

func main() {
	fmt.Println("Hello, word")
	var testValue = &value{1000, 99, 998}

	var w = bytes.NewBuffer([]byte{})
	var bw = bufio.NewWriter(w)
	var h codec.Handle = new(codec.MsgpackHandle)
	var enc = codec.NewEncoder(bw, h)

	fmt.Println(testValue.Status)
	var err = enc.Encode(testValue) //any of v1 ... v8
	bw.Flush()

	var length = len(w.Bytes())
	// var str = strconv.Itoa(length)
	fmt.Println(length)
	fmt.Println(err)

	var value2 = new(value)
	var dec = codec.NewDecoderBytes(w.Bytes(), h)
	var err2 = dec.Decode(&value2)
	fmt.Println(err2)

	fmt.Println(value2.Status)

	

}
