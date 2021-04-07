package confy

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

// A Confy is a stupid simple configuration store.
type Confy interface {
	Get(key string, ptr interface{}) error
	Set(key string, val interface{}) error
	Del(key string) error
	Keys() ([]string, error)
}

func keySplit(r rune) bool {
	return !(('a' <= r && r <= 'z') || // [a-z]
		('A' <= r && r <= 'Z') || // [A-Z]
		('0' <= r && r <= '9')) // [0-9]
}

// NormKey returns the normalised form of a given key.
//
// Confy implementations should use this to achieve consistent behaviour.
func NormKey(key string) string {
	// example - "Pokémon %% GO!"
	key = norm.NFKD.String(key) // strip accents - "Pokemon %% GO!"
	key = strings.ToLower(key)  // all lowercase - "pokemon %% go!"
	key = strings.Join(
		strings.FieldsFunc(key, keySplit), // split on non-alphanumeric chars - ["pokemon", "go"]
		"_",
	) // rejoin with underscore - "pokemon_go"

	return key
}
