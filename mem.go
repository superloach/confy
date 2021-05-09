package confy

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

var _ Confy = new(Mem)

// Mem is an in-memory Confy that uses a bytes.Buffer.
type Mem struct {
	buf bytes.Buffer
	mu  sync.Mutex
}

// Load implements Confy using io.ReadAll.
func (m *Mem) Load() ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := io.ReadAll(&m.buf)
	if err != nil {
		return nil, fmt.Errorf("readall: %w", err)
	}

	return data, nil
}

// Store implements Confy using (*bytes.Buffer).Write.
func (m *Mem) Store(data []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.buf.Reset()

	if _, err := m.buf.Write(data); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}
