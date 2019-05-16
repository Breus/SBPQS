/*
 * For this Lamport OTS implementation, we assume siging a 32-byte digest of the arbitrary length message
 * The wanted security level is 128 bit, thus for the random variables, we also pick 32-byte random variables
 */

package lamport

import (
	"crypto/rand"
	"log"
)

type keyPair struct {
	privateKey privateKey
	publicKey  publicKey
}

type privateKey struct {
	zeroVals [256][32]byte
	oneVals  [256][32]byte
}

type publicKey struct {
	hashedZeroVals [256][32]byte
	hashedOneVals  [256][32]byte
}

func KeyGen() (keypair keyPair, err error) {
	for i := 0; i < 256; i++ {
		_, err = rand.Read(keypair.privateKey.zeroVals[i][:])
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
