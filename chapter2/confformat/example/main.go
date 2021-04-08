package main

import (
	"github.com/ravitri/go-cookbook/chapter2/confformat"
)

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}

}
