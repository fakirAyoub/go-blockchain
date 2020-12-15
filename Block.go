package main

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

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

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func hash(data []byte) Hash {
	return sha1.Sum(data)
}

func (b Block) hash() Hash {
	return hash([]byte(b)[:])
}

func NewBlock(
	Index int, PreviousHash string, Data string, Timestamp time.Time,
	Difficulty int, Hash string, Nonce int, TransactionsCount int, Transactions []Transaction) *Block {
	b := new(Block)
	b.Index = Index
	b.PreviousHash = PreviousHash
	b.Data = Data
	b.Timestamp = time.Now()
	b.Difficulty = Difficulty
	b.Hash = Hash
	b.Nonce = 0
	b.TransactionsCount = TransactionsCount
	b.Transactions = Transactions
	return b
}
