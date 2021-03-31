package confy

import (
	"encoding/json"
	"fmt"
	"io"
)

// FormJSON is a JSON-based Form, ideal for structured data.
type FormJSON struct {
	// Marshal options
	EscapeHTML bool
	Prefix     string
	Indent     string

	// Unmarshal options
	DisallowUnknownFields bool
	UseNumber             bool
}

// Marshal encodes the given value into the given Writer - in this case, using json.NewEncoder.
func (f FormJSON) Marshal(w io.Writer, value interface{}) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(f.EscapeHTML)
	e.SetIndent(f.Prefix, f.Indent)

	if err := e.Encode(value); err != nil {
		return fmt.Errorf("json encode %v: %w", value, err)
	}

	return nil
}

// Unmarshal decodes the given Reader into the given pointer - in this case, using json.NewDecoder.
func (f FormJSON) Unmarshal(r io.Reader, value interface{}) error {
	d := json.NewDecoder(r)

	if f.DisallowUnknownFields {
		d.DisallowUnknownFields()
	}

	if f.UseNumber {
		d.UseNumber()
	}

	if err := d.Decode(value); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
