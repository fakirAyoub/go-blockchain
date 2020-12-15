package main

//coin
type Coin struct {
	name         string
	symbol       string
	decimals     string
	totalSupply_ int
}

func NewCoin(
	name string, symbol string, decimals string, totalSupply_ int) *Coin {
	c := new(Coin)
	c.name = name
	c.symbol = symbol
	c.decimals = decimals
	c.totalSupply_ = totalSupply_
	return c
}
