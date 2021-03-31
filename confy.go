package confy

import "fmt"

// A Confy is a stupid simple configuration store.
type Confy struct {
	Data
	Form
}

// New makes a new Confy with the given Data and Form.
func New(s Data, f Form) Confy {
	return Confy{
		Data: s,
		Form: f,
	}
}

// Load wraps a call to Data.Reader and Form.Unmarshal.
func (c Confy) Load(key string, ptr interface{}) error {
	r, err := c.Reader(key)
	if err != nil {
		return fmt.Errorf("get reader %q: %w", key, err)
	}
	defer r.Close()

	err = c.Unmarshal(r, ptr)
	if err != nil {
		return fmt.Errorf("unmarshal ptr: %w", err)
	}

	return nil
}

// Store wraps a call to Data.Writer and Form.Marshal.
func (c Confy) Store(key string, value interface{}) error {
	w, err := c.Writer(key)
	if err != nil {
		return fmt.Errorf("get writer %q: %w", key, err)
	}
	defer w.Close()

	err = c.Marshal(w, value)
	if err != nil {
		return fmt.Errorf("marshal %v: %w", value, err)
	}

	return nil
}
