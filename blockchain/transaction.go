package blockchain

import "fmt"

const REWARD int = 100

func (in *TxnInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxnOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func (txn *Transaction) IsCoinbase() bool {
	return len(txn.Inputs) == 1 && len(txn.Inputs[0].ID) == 0 && txn.Inputs[0].Out == -1
}

func CoinbaseTxn(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}

	txIn := TxnInput{
		ID:  []byte{},
		Out: -1,
		Sig: data,
	}

	txOut := TxnOutput{
		Value:  REWARD,
		PubKey: toAddress,
	}

	return &Transaction{
		ID:      nil,
		Inputs:  []TxnInput{txIn},
		Outputs: []TxnOutput{txOut},
	}
}
