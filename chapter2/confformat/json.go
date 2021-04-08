package confformat

import (
	"bytes"
	"encoding/json"
)

type JSONData struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)
	return b, nil
}

func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}
