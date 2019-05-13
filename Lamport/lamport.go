package lamport

/*
* For this Lamport OTS implementation, we assume siging a 32-byte digest of the arbitrary length message
* The wanted security level is 128 bit, thus for the random variables, we also pick 32-byte random variables
 */
type KeyPair struct {
	privateKey [32 * 32]byte
	publicKey  [32 * 32]byte
}

func KeyGen() (keypair KeyPair) {
	for i := 0; i < len(message); i++ {

	}

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
