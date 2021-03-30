package confy

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// YAMLForm is a YAML-based Form, ideal for readable configuration.
type YAMLForm struct{}

// NewYAMLForm creates a usable YAMLForm.
func NewYAMLForm() Form {
	return YAMLForm{}
}

// Marshal encodes the given value into the given Writer - in this case, using yaml.NewEncoder.
func (yf YAMLForm) Marshal(w io.Writer, value interface{}) error {
	err := yaml.NewEncoder(w).Encode(value)
	if err != nil {
		return fmt.Errorf("yaml encode %v: %w", value, err)
	}

	return nil
}

func (yf YAMLForm) Unmarshal(r io.Reader, value interface{}) error {
	err := yaml.NewDecoder(r).Decode(value)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
