package test // à remplacer par main

import (
	"os"
	"time"
)

//Block

type Block struct {
	Index             int
	PreviousHash      string
	Data              string
	Timestamp         time.Time
	Difficulty        int
	Hash              string
	Nonce             int
	TransactionsCount int
	Transactions      []Transaction
}

type Blockchain struct {
	blocks []*Block
}

func CreateBlock() Block {
	var block Block

	block.Timestamp = time.Now()
	block.Nonce = 0

	return block
}

// Transaction data structure
type Transaction struct {
	From     string
	Hash     string
	Value    string
	Gas      uint64
	GasPrice uint64
	Nonce    uint64
	To       string
	Pending  bool
}

type TransferEthRequest struct {
	PrivKey string
	To      string
	Amount  int64
}

type HashResponse struct {
	Hash string
}

type BalanceResponse struct {
	Address string
	Balance string
	Symbol  string
	Units   string
}
