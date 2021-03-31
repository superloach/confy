package confy

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// FormYAML is a YAML-based Form, ideal for readable configuration.
type FormYAML struct {
	// Marshal options
	Indent int

	// Unmarshal options
	KnownFields bool
}

// Marshal encodes the given value into the given Writer - in this case, using yaml.NewEncoder.
func (f FormYAML) Marshal(w io.Writer, value interface{}) error {
	e := yaml.NewEncoder(w)
	e.SetIndent(f.Indent)

	if err := e.Encode(value); err != nil {
		return fmt.Errorf("yaml encode %v: %w", value, err)
	}

	return nil
}

// Unmarshal decodes the given Reader into the given pointer - in this case, using yaml.NewDecoder.
func (f FormYAML) Unmarshal(r io.Reader, value interface{}) error {
	d := yaml.NewDecoder(r)
	d.KnownFields(f.KnownFields)

	if err := d.Decode(value); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
