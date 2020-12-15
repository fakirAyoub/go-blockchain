package models

type Miner struct {
	MinedBlocks      			[]Block
}

func NewMiner(MinedBlocks []Block) *Miner {
	m := new(Miner)
	m.MinedBlocks = MinedBlocks
	return m
}