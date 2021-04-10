package RSA

import (
	"crypto/rand"
	"math/big"
)

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	N *big.Int
	D *big.Int
}

func GenerateRSAKey() (*PublicKey, *PrivateKey, error) {
	p, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	q, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}

	for p.Cmp(q) == 0 {
		p, err = rand.Prime(rand.Reader, 1024)
		if err != nil {
			return nil, nil, err
		}
	}

	n := &big.Int{}
	n = n.Mul(p, q)

	EulerN := &big.Int{}
	EulerN = EulerN.Mul(p.Sub(p, bigOne), q.Sub(q, bigOne))
	E := big.NewInt(65537)

	_, x, _ := ExtendedGCD(E, EulerN)
	d := x.Add(x, EulerN)
	publicKey := &PublicKey{N: n, E: E}
	privateKey := &PrivateKey{N: n, D: d}

	return publicKey, privateKey, nil
}

func ExtendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	var value *big.Int = &big.Int{}
	if b.Cmp(bigZero) == 0 {
		return a, bigOne, bigZero
	} else {
		gcd, x, y := ExtendedGCD(b, value.Mod(a, b))
		return gcd, y, x.Sub(x, (value.Div(a, b).Mul(value, y)))
	}
}

func Encrypt(message *big.Int, publicKey *PublicKey) *big.Int {
	return big.NewInt(0).Exp(message, publicKey.E, publicKey.N)
}

func Decrypt(message *big.Int, privateKey *PrivateKey) *big.Int {
	return big.NewInt(0).Exp(message, privateKey.D, privateKey.N)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
