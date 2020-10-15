package writer

import (
	"bytes"

	"gopkg.in/yaml.v3"

	"github.com/buildpacks/pack/internal/commands"
)

type YAML struct {
	StructuredFormat
}

func NewYAML() commands.BuilderWriter {
	return &YAML{
		StructuredFormat: StructuredFormat{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				buf := bytes.NewBuffer(nil)
				if err := yaml.NewEncoder(buf).Encode(v); err != nil {
					return []byte{}, err
				}
				return buf.Bytes(), nil
			},
		},
	}
}
