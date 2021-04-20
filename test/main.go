package main

import (
	"log"
	"math/big"

	"github.com/davidleitw/RSA"
)

func main() {
	// generate key
	public, private, err := RSA.GenerateRSAKey()
	if err != nil {
		log.Println(err)
	}
	message := 1234548489

	secret := RSA.Encrypt(big.NewInt(int64(message)), public)
	log.Println("secret = ", secret)

	m := RSA.Decrypt(secret, private)
	log.Println("After Decrypt: ", m)

	message2 := big.NewInt(int64(9876554321))
	log.Println("secret = ", message2.Exp(message2, private.D, private.N))

	m = big.NewInt(0).Exp(message2, public.E, public.N)
	log.Println("After Decrypt: ", m)
}
