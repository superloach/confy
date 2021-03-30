package confy

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONForm is a JSON-based Form, ideal for structured data.
type JSONForm struct{}

// NewJSONForm creates a usable JSONForm.
func NewJSONForm() Form {
	return JSONForm{}
}

// Marshal encodes the given value into the given Writer - in this case, using json.NewEncoder.
func (jf JSONForm) Marshal(w io.Writer, value interface{}) error {
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		return fmt.Errorf("json encode %v: %w", value, err)
	}

	return nil
}

// Unmarshal decodes the given Reader into the given pointer - in this case, using json.NewDecoder.
func (jf JSONForm) Unmarshal(r io.Reader, value interface{}) error {
	err := json.NewDecoder(r).Decode(value)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
