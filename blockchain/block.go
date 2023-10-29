package blockchain

// createBlock generates a new block
func createBlock(txns []*Transactions, prevHash []byte) *Block {
	block := &Block{
		Hash:         []byte{},
		Transactions: txns,
		PrevHash:     prevHash,
		Nonce:        0,
	}

	pow := CreateProofOfWork(block)
	nonce, hash := RunProofOfWork(pow)

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
