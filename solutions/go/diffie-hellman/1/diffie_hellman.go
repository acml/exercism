// Package diffiehellman provides Diffie Hellman key exchange functions.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a private key, greater than 1 and less than p.
func PrivateKey(p *big.Int) *big.Int {
	n, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	if err != nil {
		return big.NewInt(0)
	}
	return n.Add(n, big.NewInt(2))
}

// PublicKey calculates a public key.
func PublicKey(a, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), a, p)
}

// SecretKey calculates secret key.
func SecretKey(a, B, p *big.Int) *big.Int {
	return new(big.Int).Exp(B, a, p)
}

// NewPair calculates a pair of public and private key.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	return a, PublicKey(a, p, g)
}
