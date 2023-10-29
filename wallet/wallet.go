package wallet

import (
	"blockchain-cli/blockchain"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
)

const (
	checksumLength int  = 4
	version        byte = byte(0x00)
)

// generates the address of a  wallet
func (w *Wallet) address() []byte {
	pubHash := publicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)

	checksum := checksum(versionedHash)
	finalHash := append(versionedHash, checksum...)

	return base58Encode(finalHash)
}

// creates a new wallet
func makeWallet(alias string) *Wallet {
	private, public := newPairKey()
	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
		Alias:      alias,
	}
}

// generates a new public/private key pair
func newPairKey() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panicln("Error generating private key.")
	}

	pub := append(private.PublicKey.X.Bytes(), private.Y.Bytes()...)

	return *private, pub
}

// returns the hash of a given public key
func publicKeyHash(publicKey []byte) []byte {
	hashedPublicKey := sha256.Sum256(publicKey)

	hasher := sha256.New()
	_, err := hasher.Write(hashedPublicKey[:])
	if err != nil {
		log.Panicln(err)
	}

	return hasher.Sum(nil)
}

func checksum(ripeMdHash []byte) []byte {
	firsHash := sha256.Sum256(ripeMdHash)
	secondHash := sha256.Sum256(firsHash[:])

	return secondHash[:checksumLength]
}

func GetBalance(address string) int {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Database.Close()

	balance := 0
	UTXOs := chain.FindUTXOs(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}
