package confy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// DataFile is a filesystem-based Data.
type DataFile struct {
	Base string
	Ext  string
}

// NewDataFile creates a DataFile with a given base directory and file extension.
func NewDataFile(base, ext string) (Data, error) {
	abase, err := filepath.Abs(base)
	if err != nil {
		return nil, fmt.Errorf("abs %q: %w", base, err)
	}

	return DataFile{
		Base: abase,
		Ext:  ext,
	}, nil
}

// ConfigDataFile creates a DataFile with the given app name and file extension.
func ConfigDataFile(app, ext string) (Data, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("get config dir: %w", err)
	}

	return NewDataFile(filepath.Join(dir, app), ext)
}

// fp creates the filepath of a given key, stripping characters as necessary.
func (d DataFile) fp(key string) string {
	return filepath.Join(
		d.Base,
		strings.Join(
			strings.FieldsFunc(
				strings.ToLower(
					norm.NFKD.String(key), // strip accents
				), // only lowercase
				func(r rune) bool {
					return r < 'a' || r > 'z'
				}, // delimit on non-alpha
			), // split into chunks of alpha
			"_", // rejoin with underscore
		)+d.Ext, // add the extension
	) // join onto base
}

// Reader fetches a ReadCloser for the given key - in this case, using os.Open.
func (d DataFile) Reader(key string) (io.ReadCloser, error) {
	fp := d.fp(key)

	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", fp, err)
	}

	return f, nil
}

// Writer fetches a WriteCloser for the given key - in this case, using os.Create.
func (d DataFile) Writer(key string) (io.WriteCloser, error) {
	fp := d.fp(key)

	f, err := os.Create(fp)
	if err != nil {
		return nil, fmt.Errorf("create %q: %w", fp, err)
	}

	return f, nil
}

// Delete removes the given key if it exists - in this case, using os.Remove.
func (d DataFile) Delete(key string) error {
	fp := d.fp(key)

	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("remove %q: %w", fp, err)
	}

	return nil
}

// Keys returns available keys - in this case, using os.ReadDir.
func (d DataFile) Keys() ([]string, error) {
	dirents, err := os.ReadDir(d.Base)
	if err != nil {
		return nil, fmt.Errorf("read dir %q: %w", d.Base, err)
	}

	ks := make([]string, 0, len(dirents))

	for _, dirent := range dirents {
		if !dirent.Type().IsRegular() {
			continue
		}

		n := dirent.Name()
		if strings.HasSuffix(n, d.Ext) {
			ks = append(ks, strings.TrimSuffix(n, d.Ext))
		}
	}

	return ks, nil
}
