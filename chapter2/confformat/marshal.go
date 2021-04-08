package confformat

import "fmt"

func MarshalAll() error {

	j := JSONData{
		Name: "RaviJSON",
		Age:  32,
	}

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	y := YAMLData{
		Name: "RaviYAML",
		Age:  32,
	}

	yamlRes, err := y.ToYaml()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal=", yamlRes.String())

	return nil
}
