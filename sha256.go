package kalkan

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type sha256Hasher struct{ hash.Hash }

func (h sha256Hasher) WriteString(s string) (string, error) {
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewSHA256 returns a new SHA256 Hasher.
func NewSHA256() Hasher {
	return sha1Hasher{sha256.New()}
}
