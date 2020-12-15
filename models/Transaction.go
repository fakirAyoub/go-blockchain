package models

import "time"

type Transaction struct {
	TxId               string
	SourceAddress      string
	DestinationAddress string
	Amount             int64
	TransactionDate	   time.Time
}

func NewTransaction(TxId string, SourceAddress string, DestinationAddress string, Amount int64,
	TransactionDate time.Time) *Transaction {
	tx := new(Transaction)
	tx.TxId = TxId
	tx.SourceAddress = SourceAddress
	tx.DestinationAddress = DestinationAddress
	tx.Amount = Amount
	tx.TransactionDate = TransactionDate
	return tx
}

func (t *Transaction) calculateUTXO(senderAddress string) {
	//Public Address A:
	//N inputs and M outputs
	//N - M should be greater than 0
}