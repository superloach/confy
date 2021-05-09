package confy

import "fmt"

// Func is a Confy that just contains Load and Store funcs.
type Func struct {
	LoadFn  func() ([]byte, error)
	StoreFn func([]byte) error
}

// Load implements Confy using LoadFn.
func (f Func) Load() ([]byte, error) {
	data, err := f.LoadFn()
	if err != nil {
		return nil, fmt.Errorf("loadfn: %w", err)
	}

	return data, nil
}

// Store implements Confy using StoreFn.
func (f Func) Store(data []byte) error {
	if err := f.StoreFn(data); err != nil {
		return fmt.Errorf("storefn: %w", err)
	}

	return nil
}
