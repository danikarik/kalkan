package kalkan

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

type sha1Hasher struct{ hash.Hash }

func (h sha1Hasher) WriteString(s string) (string, error) {
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewSHA1 returns a new SHA1 Hasher.
func NewSHA1() Hasher {
	return sha1Hasher{sha1.New()}
}
