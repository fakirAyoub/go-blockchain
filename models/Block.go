package models

import "time"

type Block struct {
	PreviousHash      			string
	CreationTimestamp 			time.Time
	Difficulty        			int
	MerkleRootHash              string
	TransactionsCount 			int
	MinedBy      	  			Miner
	Transactions      			[]Transaction
}

func NewBlock(
	PreviousHash string,
	Difficulty int, MerkleRootHash string, TransactionsCount int, Transactions []Transaction) *Block {
	b := new(Block)
	b.PreviousHash = PreviousHash
	b.CreationTimestamp = time.Now()
	b.Difficulty = Difficulty
	b.MerkleRootHash = MerkleRootHash
	b.TransactionsCount = TransactionsCount
	b.Transactions = Transactions
	return b
}