package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

func CreateBlockChain(address string) *BlockChain {
	log.Println("Blockchain already exists")
	var lastHash []byte

	if BlockChainExists() {
		log.Println("Blockchain already exists")
		runtime.Goexit()
	}

	ops := badger.DefaultOptions(DB_PATH)
	ops.Logger = nil

	db, err := badger.Open(ops)
	handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		log.Printf("ADDRESS: %s", address)
		return err
	})
	handle(err)

	chain := BlockChain{
		LastHash: lastHash,
		Database: db,
	}

	return &chain
}

func BlockChainExists() bool {
	_, err := os.Stat(DB_FILE)
	return !os.IsNotExist(err)
}

func (b *Block) hashTransactions() []byte {
	var txnHashes [][]byte
	var txnHash [32]byte

	for _, txn := range b.Transactions {
		txnHashes = append(txnHashes, txn.ID)
	}

	txnHash = sha256.Sum256(bytes.Join(txnHashes, []byte{}))
	return txnHash[:]
}
