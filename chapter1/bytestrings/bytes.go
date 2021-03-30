package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func WorkWithBuffer() error {
	rawString := "This is Ravi Trivedi"
	b := Buffer(rawString)
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
	return nil
}
