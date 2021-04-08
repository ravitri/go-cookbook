package confformat

import (
	"bytes"
	"gopkg.in/yaml.v2"
)

type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func (t *YAMLData) ToYaml() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)
	return b, nil

}

func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
