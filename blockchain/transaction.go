package blockchain

func (in *TxnInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxnOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (txn *Transaction) IsCoinbase() bool {
	return len(txn.Inputs) == 1 && len(txn.Inputs[0].ID) == 0 && txn.Inputs[0].Out == -1
}
