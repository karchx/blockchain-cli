package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

func CreateBlockChain(address string) *BlockChain {
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

func ContinueBlockChain(address string) *BlockChain {
	if !BlockChainExists() {
		fmt.Println("No blockchain found, please a create one first")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions(DB_PATH)
	opts.Logger = nil
	db, err := badger.Open(opts)
	handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = nil
			return nil
		})

		handle(err)
		return err
	})
	handle(err)

	chain := BlockChain{
		lastHash,
		db,
	}
	return &chain
}

func BlockChainExists() bool {
	_, err := os.Stat(DB_FILE)
	return !os.IsNotExist(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iterator := BlockChainIterator{
		CurrentHash: chain.LastHash,
		Database:    chain.Database,
	}

	return &iterator
}

func (it *BlockChainIterator) Next() *Block {
	var b *Block

	err := it.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(it.CurrentHash)
		handle(err)

		err = item.Value(func(val []byte) error {
			b = deserialize(val)
			return nil
		})
		handle(err)
		return err
	})
	handle(err)

	it.CurrentHash = b.PrevHash
	return b
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

func (chain *BlockChain) findUnspentTxns(address string) []Transaction {
	var uxtos []Transaction

	spentTXNs := make(map[string][]int)
	it := chain.Iterator()

	for {
		block := it.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIDx, out := range tx.Outputs {
				if spentTXNs[txID] != nil {
					for _, spentOut := range spentTXNs[txID] {
						if spentOut == outIDx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlocked(address) {
					uxtos = append(uxtos, *tx)
				}
			}

			if !tx.IsCoinbase() {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxnID := hex.EncodeToString(in.ID)
						spentTXNs[inTxnID] = append(spentTXNs[inTxnID], in.Out)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return uxtos
}

func (chain *BlockChain) FindUTXOs(address string) []TxnOutput {
	var UTXOs []TxnOutput
	unspent := chain.findUnspentTxns(address)

	for _, tx := range unspent {
		for _, out := range tx.Outputs {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}
