package kalkan

import (
	"encoding/hex"
	"hash"

	"github.com/martinlindhe/gogost/gost341194"
)

type gost34311Hasher struct{ hash.Hash }

func (h gost34311Hasher) WriteString(s string) (string, error) {
	if _, err := h.Write([]byte(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewGOST returns a new GOST34311 Hasher.
func NewGOST() Hasher {
	return sha1Hasher{gost341194.New(gost341194.SboxDefault)}
}
