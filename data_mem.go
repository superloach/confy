package confy

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

// ErrMemNoKey occurs when a given key does not exist.
var ErrMemNoKey = errors.New("no such key in MemData")

type memBuffer struct {
	bytes.Buffer
}

func (mb *memBuffer) Close() error {
	return nil
}

// MemData is a concurrent-safe, in-memory Data.
type MemData struct {
	ma map[string]*memBuffer
	mu *sync.RWMutex
}

// NewMemData creates a usable MemData.
func NewMemData() Data {
	return MemData{
		ma: make(map[string]*memBuffer),
		mu: &sync.RWMutex{},
	}
}

// Reader fetches a ReadCloser for the given key - in this case, using internal maps and locks.
func (ms MemData) Reader(key string) (io.ReadCloser, error) {
	ms.mu.RLock()
	buf, ok := ms.ma[key]
	ms.mu.RUnlock()

	if !ok {
		return nil, ErrMemNoKey
	}

	return buf, nil
}

// Writer fetches a WriteCloser for the given key - in this case, using internal maps and locks.
func (ms MemData) Writer(key string) (io.WriteCloser, error) {
	ms.mu.RLock()
	buf, ok := ms.ma[key]
	ms.mu.RUnlock()

	if !ok {
		buf := &memBuffer{
			Buffer: bytes.Buffer{},
		}

		ms.mu.Lock()
		ms.ma[key] = buf
		ms.mu.Unlock()
	}

	return buf, nil
}

// Delete removes the given key if it exists - in this case, using internal maps and locks.
func (ms MemData) Delete(key string) error {
	ms.mu.Lock()
	delete(ms.ma, key)
	ms.mu.Unlock()

	return nil
}

// Keys returns available keys - in this case, using internal maps and locks.
func (ms MemData) Keys() ([]string, error) {
	// lock this whole section to ensure len doesn't change (for minimal alloc)
	ms.mu.RLock()
	ks := make([]string, 0, len(ms.ma))

	for k := range ms.ma {
		ks = append(ks, k)
	}
	ms.mu.RUnlock()

	return ks, nil
}
