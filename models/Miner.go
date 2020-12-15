package models

type Miner struct {
	minedBlocks      			[]Block
	reward						int
}

func NewMiner(minedBlocks []Block) *Miner {
	m := new(Miner)
	m.minedBlocks = minedBlocks
	return m
}

func (m *Miner) getReward(reward int) {
	m.reward = m.reward + reward
}

func (m *Miner) validateTransaction(transaction Transaction) bool {
	return true
}

func (m *Miner) createBlock(transactions []Transaction) bool {
	return true
}


