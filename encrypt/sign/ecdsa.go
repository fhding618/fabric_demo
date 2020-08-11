package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func ECDSADemo()  {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(err)
	}
	publicKey := &privateKey.PublicKey

	fmt.Printf("%d\n", rand.Reader)
	fmt.Printf("%v\n", publicKey)

	message := []byte("ECDSA Demo")

	// Sign
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, message)

	fmt.Printf("%v\n", r)

	fmt.Printf("%v\n", s)

	// Verify
	flag := ecdsa.Verify(publicKey, []byte("ab"), r, s)
	fmt.Printf("%v\n", flag)

}