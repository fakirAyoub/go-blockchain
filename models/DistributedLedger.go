package models

type DistributedLedger struct {
	genesis Block
	blocks []Block
	totalSupply int
}

func NewDistributedLedger(genesis Block, blocks []Block, totalSupply int) *DistributedLedger {
	d := new (DistributedLedger)
	d.genesis = genesis
	d.blocks = blocks
	d.totalSupply = totalSupply
	return d
}