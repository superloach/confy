package confy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// OS is a Confy which uses functionality from package os.
type OS struct {
	Directory string
	Extension string
}

// NewOS creates a new OS with the given directory and extension.
func NewOS(dir, ext string) (*OS, error) {
	adir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("abs %q: %w", dir, err)
	}

	return &OS{
		Directory: adir,
		Extension: ext,
	}, nil
}

func (o *OS) fp(key string) string {
	return filepath.Join(o.Directory, NormKey(key)+o.Extension)
}

// Get loads JSON data into ptr from the file corresponding to the given key.
func (o *OS) Get(key string, ptr interface{}) error {
	fp := o.fp(key)

	f, err := os.Open(fp)
	if err != nil {
		return fmt.Errorf("os open %q: %w", fp, err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(ptr)
	if err != nil {
		return fmt.Errorf("json decode ptr: %w", err)
	}

	return nil
}

// Set writes val in JSON format to the file corresponding to the given key.
func (o *OS) Set(key string, val interface{}) error {
	fp := o.fp(key)

	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("os create %q: %w", fp, err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(val)
	if err != nil {
		return fmt.Errorf("json encode %v: %w", val, err)
	}

	return nil
}

// Del removes the file corresponding to the given key.
func (o *OS) Del(key string) error {
	fp := o.fp(key)

	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("os remove %q: %w", fp, err)
	}

	return nil
}

// Keys fetches a list of existing keys using os.ReadDir, filtered by o.Extension and NormKey.
func (o *OS) Keys() ([]string, error) {
	ents, err := os.ReadDir(o.Directory)
	if err != nil {
		return nil, fmt.Errorf("os read dir %q: %w", o.Directory, err)
	}

	ks := make([]string, 0, len(ents))

	for _, ent := range ents {
		name := ent.Name()
		if !strings.HasSuffix(name, o.Extension) {
			continue
		}

		k := strings.TrimSuffix(name, o.Extension)
		if k != NormKey(k) {
			continue
		}

		ks = append(ks, name)
	}

	return ks, nil
}
