package main

import (
	"log"
	"math/big"

	"github.com/davidleitw/RSA"
)

func main() {
	public, private, err := RSA.GenerateRSAKey()
	if err != nil {
		log.Println(err)
	}
	message := 1234548489

	secret := RSA.Encrypt(big.NewInt(int64(message)), public)
	log.Println("secret = ", secret)

	m := RSA.Decrypt(secret, private)
	log.Println("Afert Decrypt: ", m)
}
