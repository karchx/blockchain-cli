package blockchain

import (
	"math/big"

	"github.com/dgraph-io/badger"
)

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

type Transaction struct {
	ID      []byte
	Inputs  []TxnInput
	Outputs []TxnOutput
}

type TxnInput struct {
	ID  []byte
	Out int
	Sig string
}

type TxnOutput struct {
	Value  int
	PubKey string
}
