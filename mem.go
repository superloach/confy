package confy

import (
	"encoding/json"
	"fmt"
	"sync"
)

var _ Confy = &Mem{}

type Mem struct {
	mu sync.RWMutex
	ma map[string][]byte
}

func NewMem() *Mem {
	return &Mem{
		mu: sync.RWMutex{},
		ma: map[string][]byte{},
	}
}

func (m *Mem) Get(key string, ptr interface{}) error {
	key = NormKey(key)

	m.mu.RLock()
	bs := m.ma[key]
	m.mu.RUnlock()

	err := json.Unmarshal(bs, ptr)
	if err != nil {
		return fmt.Errorf("json unmarshal %q: %w", bs, err)
	}

	return nil
}

func (m *Mem) Set(key string, val interface{}) error {
	key = NormKey(key)

	bs, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("json marshal %v: %w", val, err)
	}

	m.mu.Lock()
	m.ma[key] = bs
	m.mu.Unlock()

	return nil
}

func (m *Mem) Del(key string) error {
	key = NormKey(key)

	m.mu.Lock()
	delete(m.ma, key)
	m.mu.Unlock()

	return nil
}

func (m *Mem) Keys() ([]string, error) {
	m.mu.RLock()

	ks := make([]string, 0, len(m.ma))

	for k := range m.ma {
		ks = append(ks, k)
	}

	m.mu.RUnlock()

	return ks, nil
}
