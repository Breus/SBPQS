package wots

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"math"
)

//Scheme consists of the parameters for the scheme
type Scheme struct {
	n       int // Security parameter in bytes
	w       int // Winternitz parameter
	logw    int // Winternitz parameter log2
	length1 int // Length paramter 1
	length2 int // Length paramter 2
	length  int // Length paramter (1+2)
}

// InitScheme initializes the scheme parameters
func InitScheme(n int, w int) (scheme *Scheme, err error) {
	if (w != 0) && (w&(w-1)) != 0 {
		return nil, errors.New("w should be of power 2 and larger than 0")
	}
	if n != 32 && n != 64 {
		return nil, errors.New("The scheme supports only n = 32, or n = 64")
	}
	scheme.n = n
	scheme.w = w
	scheme.logw = int(math.Log2(float64(w)))
	scheme.length1 = int(math.Ceil(float64((n * 8) / scheme.logw)))
	scheme.length2 = int(math.Floor(math.Log2(float64(scheme.length1*(scheme.w-1))) /
		float64(scheme.logw)))
	scheme.length = scheme.length1 + scheme.length2

	return scheme, err
}

//generateSeed generates a n-byte private seed using the golang random function
func (s *Scheme) generateSeed() (seed []byte, err error) {
	_, err = rand.Read(seed[0:s.n])
	return
}

//expandKey expands a random n-byte seed to a n*length signing key
func (s *Scheme) expandKey(seed []byte) (expandedKey []byte) {
	return
}

//genVerKey genenerates the corresponding verification key

//hash generates a n-byte digest of the input
func (s *Scheme) hash(input []byte) (digest []byte, err error) {
	switch s.n {
	case 32:
		digest := sha256.Sum256(input)
		return digest[:], nil
	case 64:
		digest := sha512.Sum512(input)
		return digest[:], nil
	default:
		return nil, errors.New("The scheme supports either n = 32 or n = 64")
	}
}
