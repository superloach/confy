package confy

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONForm is a JSON-based Form, ideal for structured data.
type JSONForm struct {
	// Marshal options
	EscapeHTML bool
	Prefix     string
	Indent     string

	// Unmarshal options
	DisallowUnknownFields bool
	UseNumber             bool
}

// Marshal encodes the given value into the given Writer - in this case, using json.NewEncoder.
func (jf JSONForm) Marshal(w io.Writer, value interface{}) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(jf.EscapeHTML)
	e.SetIndent(jf.Prefix, jf.Indent)

	if err := e.Encode(value); err != nil {
		return fmt.Errorf("json encode %v: %w", value, err)
	}

	return nil
}

// Unmarshal decodes the given Reader into the given pointer - in this case, using json.NewDecoder.
func (jf JSONForm) Unmarshal(r io.Reader, value interface{}) error {
	d := json.NewDecoder(r)

	if jf.DisallowUnknownFields {
		d.DisallowUnknownFields()
	}

	if jf.UseNumber {
		d.UseNumber()
	}

	if err := d.Decode(value); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
