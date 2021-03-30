package confy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// FileData is a filesystem-based Data.
type FileData struct {
	Base string
	Ext  string
}

// NewFileData creates a FileData with a given base directory and file extension.
func NewFileData(base, ext string) (Data, error) {
	abase, err := filepath.Abs(base)
	if err != nil {
		return nil, fmt.Errorf("abs %q: %w", base, err)
	}

	return FileData{
		Base: abase,
		Ext:  ext,
	}, nil
}

// ConfigFileData creates a FileData with the given app name and file extension.
func ConfigFileData(app, ext string) (Data, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("get config dir: %w", err)
	}

	return NewFileData(filepath.Join(dir, app), ext)
}

// fp creates the filepath of a given key, stripping characters as necessary.
func (fs FileData) fp(key string) string {
	return filepath.Join(
		fs.Base,
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
		)+fs.Ext, // add the extension
	) // join onto base
}

// Reader fetches a ReadCloser for the given key - in this case, using os.Open.
func (fs FileData) Reader(key string) (io.ReadCloser, error) {
	fp := fs.fp(key)

	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", fp, err)
	}

	return f, nil
}

// Writer fetches a WriteCloser for the given key - in this case, using os.Create.
func (fs FileData) Writer(key string) (io.WriteCloser, error) {
	fp := fs.fp(key)

	f, err := os.Create(fp)
	if err != nil {
		return nil, fmt.Errorf("create %q: %w", fp, err)
	}

	return f, nil
}

// Delete removes the given key if it exists - in this case, using os.Remove.
func (fs FileData) Delete(key string) error {
	fp := fs.fp(key)

	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("remove %q: %w", fp, err)
	}

	return nil
}

// Keys returns available keys - in this case, using os.ReadDir.
func (fs FileData) Keys() ([]string, error) {
	ds, err := os.ReadDir(fs.Base)
	if err != nil {
		return nil, fmt.Errorf("read dir %q: %w", fs.Base, err)
	}

	ks := make([]string, 0, len(ds))

	for _, d := range ds {
		if !d.Type().IsRegular() {
			continue
		}

		n := d.Name()
		if strings.HasSuffix(n, fs.Ext) {
			ks = append(ks, strings.TrimSuffix(n, fs.Ext))
		}
	}

	return ks, nil
}
