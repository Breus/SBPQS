package lamport

import (
	"crypto/rand"
	"log"
)

/*
* For this Lamport OTS implementation, we assume siging a 32-byte digest of the arbitrary length message
* The wanted security level is 128 bit, thus for the random variables, we also pick 32-byte random variables
 */
type KeyPair struct {
	privateKey privateKey
	publicKey  publicKey
}

type privateKey struct {
	zero_vals [256][32]byte
	one_vals  [256][32]byte
}

type publicKey struct {
	h_zero_vals [256][32]byte
	h_one_vals  [256][32]byte
}

func KeyGen() (keypair KeyPair, er error) {
	for i := 0; i < 256; i++ {
		_, err := rand.Read(keypair.privateKey.zero_vals[i][:])
		if err != nil {
			log.Fatal("Random reader crashed in key generation")
			return
		}
	}

	return

}

func Sign(message []byte) (signature []byte) {
	for i := 0; i < len(message); i++ {
		for j := 0; j < 8; j++ {

		}
	}
	return signature
}

func Verify(message []byte, signature []byte) (accept bool) {
	return true
}
