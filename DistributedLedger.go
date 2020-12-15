package main

// merkle root, liste transactions et le root du bloc précédent

type DistributedLedger struct {
	merkleRoot   string
	transactions Transaction
	root         Block
}

func NewDistributedLedger(
	merkleRoot string, transactions Transaction, root Block) *DistributedLedger {
	dl := new(DistributedLedger)
	dl.merkleRoot = merkleRoot
	dl.transactions = transactions
	dl.root = root
	return dl
}
