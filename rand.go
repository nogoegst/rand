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
	"math/big"
	mrand "math/rand"
)

// Reader is crypto/rand.Reader
var Reader = crand.Reader

// Rand is crypto/rand.Read
func Read(b []byte) (n int, err error) {
	return io.ReadFull(Reader, b)
}

// BigInt is crypto/rand.Int
func BigInt(rand io.Reader, max *big.Int) (n *big.Int, err error) {
	return crand.Int(rand, max)
}

// Prime is crypto/rand.Prime
func Prime(rand io.Reader, bits int) (p *big.Int, err error) {
	return Prime(rand, bits)
}

/**/

// CryptoRandomSource implements math/rand.Source interface using
// cryptographically secure source (crypto/rand).
type CryptoRandomSource struct {
	reader io.Reader
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
	_, err := io.ReadFull(s.reader, buf)
	if err != nil {
		panic(err)
	}
	return binary.BigEndian.Uint64(buf)
}

func NewSource(reader io.Reader) mrand.Source {
	return &CryptoRandomSource{
		reader: reader,
	}
}

func New() *mrand.Rand {
	return mrand.New(NewSource(crand.Reader))
}

func NewWithReader(reader io.Reader) *mrand.Rand {
	return mrand.New(NewSource(reader))
}
