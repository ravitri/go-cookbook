package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {

	// This can also be done as following as an alternative:
	// 1. rawBytes := []byte(rawString)
	//    var b = new(bytes.Buffer)
	//    b.Write(rawBytes)
	// 2. b = bytes.NewBuffer(rawBytes)
	b := bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return "", err
	}
	return string(b), nil
}
