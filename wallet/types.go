package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

// Wallet a wallet consists of a public/private key pair
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Alias      string
}

// Wallets each user may have many wallets for privacy
type Wallets struct {
	Wallets map[string]*Wallet
}

type P256Curve struct {
	elliptic.Curve
}

func (p P256Curve) GobEncode() ([]byte, error) {
	return []byte("P256"), nil
}

func (p *P256Curve) GobDecode(data []byte) error {
	if string(data) != "P256" {
		return nil
	}

	p.Curve = elliptic.P256()
	return nil
}
