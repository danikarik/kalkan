package kalkan

import "hash"

// Hasher extends standard hash.Hash interface.
type Hasher interface {
	hash.Hash

	WriteString(s string) (string, error)
}
