package wots

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"math"
)

// Parameters hold the parameters required to compute signatures with the WOTS-T scheme
type Parameters struct {
	n       int // Security parameter in BYTES
	w       int // Winternitz parameter
	logw    int // Winternitz parameter log2
	length1 int // Length paramter 1
	length2 int // Length paramter 2
	length  int // Length paramter (1+2)

}

// InitParameters initializes the WOTS-T paramters
func InitParameters(w int, n int, msg string) (params Parameters, err error) {
	if (w & (w - 1)) != 0 {
		return params, errors.New("w should be of power 2")
	}
	if n != 32 && n != 64 {
		return params, errors.New("n should be 32 or 64 bytes security")
	}
	params.n = n
	params.w = w
	params.logw = int(math.Log2(float64(w)))
	params.length1 = int(math.Ceil(float64(8 * params.n / params.logw)))
	params.length2 = int(math.Floor(math.Log2(float64(params.length1*(params.w-1))) /
		float64(params.logw)))
	params.length = params.length1 + params.length2

	return params, nil
}

func KeyGen(seed []byte, params Parameters) {
	sk := expandKey(seed, params)

}

func expandKey(key []byte, params Parameters) (expandedKey []byte) {
	for i := 0; i < params.length; i++ {
		counter := make([]byte, params.n)
		for j := 0; i < params.n; j++ {
			counter = append(counter, byte(i))
		}
		expandedKey = append(expandedKey, prf(counter, key, params.n)...)
	}
	return
}

func prf(in []byte, key []byte, n int) (out []byte) {
	var buffer []byte
	buffer = append(buffer, key...)
	buffer = append(buffer, in...)
	return hash(buffer, n)
}

func hash(in []byte, n int) []byte {
	switch n {
	case 32:
		checksum := sha256.Sum256(in)
		return checksum[:]
	case 64:
		checksum := sha512.Sum512(in)
		return checksum[:]
	default:
		checksum := sha256.Sum256(in)
		return checksum[:]
	}
}
