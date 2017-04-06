// rand.go - cryprographically secure math/rand source.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to rand, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package rand

import (
	crand "crypto/rand"
	"encoding/binary"
	"io"
	mrand "math/rand"
)

// CryptoRandomSource implements math/rand.Source interface using
// cryptographically secure source (crypto/rand).
type CryptoRandomSource struct {
}

// seed is ignored due to there is no seed
func (s *CryptoRandomSource) Seed(seed int64) {
}

func (s *CryptoRandomSource) Int63() int64 {
	u := s.Uint64()
	return int64(u << 1 >> 1)
}

func (s *CryptoRandomSource) Uint64() uint64 {
	buf := make([]byte, 8)
	_, err := io.ReadFull(crand.Reader, buf)
	if err != nil {
		panic(err)
	}
	return binary.BigEndian.Uint64(buf)
}

func NewSource() mrand.Source {
	return &CryptoRandomSource{}
}

func New() *mrand.Rand {
	return mrand.New(NewSource())
}
