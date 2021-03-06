package main

import (
	"bytes"
	"fmt"

	"github.com/ravitri/go-cookbook/chapter1/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("Ravi Trivedi"))
	out := &bytes.Buffer{}

	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err) //TODO
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}

}
