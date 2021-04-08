package confformat

import "fmt"

const (
	exampleJSON = `{"name":"RaviJSON","age":32}`
	exampleYAML = `name: RaviYAML
age: 32
`
)

func UnmarshalAll() error {
	j := JSONData{}
	y := YAMLData{}

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal=", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("YAML Unmarshal=", y)

	return nil
}
