package confy

// A Confy is a stupid simple configuration store.
type Confy interface {
	Load() ([]byte, error)
	Store([]byte) error
}
