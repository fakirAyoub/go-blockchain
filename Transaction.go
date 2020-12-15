package main

// Transaction
type Transaction struct {
	TxId               string
	SourceAddress      string
	DestinationAddress string
	Amount             int64
	UnsignedTx         string
	SignedTx           string
}

//type TransferEthRequest struct {
//	PrivKey string
//	To      string
//	Amount  int64
//}
//
//type HashResponse struct {
//	Hash string
//}
//
//type BalanceResponse struct {
//	Address string
//	Balance string
//	Symbol  string
//	Units   string
//}

//func CreateTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
//
//}

func NewTransaction(TxId string, SourceAddress string, DestinationAddress string, Amount int64,
	UnsignedTx string, SignedTx string) *Transaction {
	tx := new(Transaction)
	tx.TxId = TxId
	tx.SourceAddress = SourceAddress
	tx.DestinationAddress = DestinationAddress
	tx.Amount = Amount
	tx.UnsignedTx = UnsignedTx
	tx.SignedTx = SignedTx
	return tx
}
