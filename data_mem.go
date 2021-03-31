package confy

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

// ErrMemNoKey occurs when a given key does not exist.
var ErrMemNoKey = errors.New("no such key in DataMem")

type memBuffer struct {
	bytes.Buffer
}

func (mb *memBuffer) Close() error {
	return nil
}

// DataMem is a concurrent-safe, in-memory Data.
type DataMem struct {
	ma map[string]*memBuffer
	mu *sync.RWMutex
}

// NewDataMem creates a usable DataMem.
func NewDataMem() DataMem {
	return DataMem{
		ma: make(map[string]*memBuffer),
		mu: &sync.RWMutex{},
	}
}

// Reader fetches a ReadCloser for the given key - in this case, using internal maps and locks.
func (d DataMem) Reader(key string) (io.ReadCloser, error) {
	d.mu.RLock()
	buf, ok := d.ma[key]
	d.mu.RUnlock()

	if !ok {
		return nil, ErrMemNoKey
	}

	return buf, nil
}

// Writer fetches a WriteCloser for the given key - in this case, using internal maps and locks.
func (d DataMem) Writer(key string) (io.WriteCloser, error) {
	d.mu.RLock()
	buf, ok := d.ma[key]
	d.mu.RUnlock()

	if !ok {
		buf := &memBuffer{
			Buffer: bytes.Buffer{},
		}

		d.mu.Lock()
		d.ma[key] = buf
		d.mu.Unlock()
	}

	return buf, nil
}

// Delete removes the given key if it exists - in this case, using internal maps and locks.
func (d DataMem) Delete(key string) error {
	d.mu.Lock()
	delete(d.ma, key)
	d.mu.Unlock()

	return nil
}

// Keys returns available keys - in this case, using internal maps and locks.
func (d DataMem) Keys() ([]string, error) {
	// lock this whole section to ensure len doesn't change (for minimal alloc)
	d.mu.RLock()
	ks := make([]string, 0, len(d.ma))

	for k := range d.ma {
		ks = append(ks, k)
	}
	d.mu.RUnlock()

	return ks, nil
}
