package wots

/*
 * This package supports n = 256-bits (32B) or n = 512=bits (64B), hence 32B or 64B message digests are signed
 */
import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"math"
)

// SchemeParameters holds the parameters defining the WOTS scheme instance
type SchemeParameters struct {
	n       int // Security parameter in bytes
	w       int // Winternitz parameter
	logw    int // Winternitz parameter log2
	length1 int // Length paramter 1
	length2 int // Length paramter 2
	length  int // Length paramter (1+2)

}

// InitParameters initializes the WOTS-T paramters
func InitParameters(w int, n int) (params SchemeParameters, err error) {
	if (w != 0) && (w&(w-1)) != 0 {
		return params, errors.New("w should be of power 2 and larger than 0")
	}
	if n != 32 && n != 64 {
		return params, errors.New("n should be 32 or 64 bytes security")
	}
	params.n = n
	params.w = w
	params.logw = int(math.Log2(float64(w)))
	params.length1 = int(math.Ceil(float64((n * 8) / params.logw)))
	params.length2 = int(math.Floor(math.Log2(float64(params.length1*(params.w-1))) /
		float64(params.logw)))
	params.length = params.length1 + params.length2

	return params, err
}

//KeyGen genreates a Winternitz One Time Signature Keypair (sk,pk)
func KeyGen(params SchemeParameters) (sk [][]byte, pk [][]byte) {
	sk = make([][]byte, params.length)
	for i := 0; i < params.length; i++ {
		for j := 0; j < params.n; j++ {
			sk[i] = append(sk[i])
		}

	}

	return
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
