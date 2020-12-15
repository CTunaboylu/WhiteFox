package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash"
	"log"
	"math/big"
	"os"
)

type Identity struct {
	private_key *ecdsa.PrivateKey
	public_key  ecdsa.PublicKey
	algorithm   string
	curve       elliptic.Curve
}

type Signature struct {
	signature []uint8
	r         *big.Int
	s         *big.Int
	message   string // change this with Transaction
	// owner     User
}

func Form_Identity() Identity {
	id := Identity{}
	id.algorithm = "ECDSA-P256"
	id.curve = elliptic.P256()
	privkey, err := ecdsa.GenerateKey(id.curve, rand.Reader)
	if err != nil {
		fmt.Println("EC key generation failed")
	} else {
		id.private_key = privkey
	}
	id.public_key = id.private_key.PublicKey

	return id

}

// func (id *Identity) Sign(message string) ([]uint8, *big.Int, *big.Int) {
func (id *Identity) Sign(message string) Signature {
	var h hash.Hash
	h = sha256.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	// io.WriteString(h, message)
	h.Write([]byte(message))
	sign_hash := h.Sum(nil)
	r, s, serr := ecdsa.Sign(rand.Reader, id.private_key, sign_hash)

	if serr != nil {
		fmt.Println(serr)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(sign_hash, s.Bytes()...)

	// fmt.Printf("Signature before injecting > value: %v bytes: %b , with the type %T\n", signature, signature, signature)

	sig := Signature{
		signature: signature,
		r:         r,
		s:         s,
		message:   message,
	}
	return sig

}

// func (id *Identity) Verify_Signature(s_hash []uint8, r *big.Int, s *big.Int) bool {
func (id *Identity) Verify_Signature(s *Signature) bool {
	verify_status := ecdsa.Verify(&id.public_key, s.signature, s.r, s.s)
	fmt.Printf("Hashes match: %v\n", verify_status) // should be true
	log.Print("Hashes match: ", verify_status)      // should be true
	// fmt.Printf("signature: %v, r: %v, s: %v \n", s_hash, r, s)
	return verify_status
}

func main() {

	id := Form_Identity()
	signature := id.Sign("Hello World!")
	// signature, r, s := id.Sign("Hello World!")
	// result := id.Verify_Signature(signature, r, s)
	result := id.Verify_Signature(&signature)
	fmt.Println(result)
}
