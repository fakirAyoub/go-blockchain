package models

type Coin struct {
	name         string
	symbol       string
	decimals     string
}

func NewCoin(
	name string, symbol string, decimals string) *Coin {
	c := new(Coin)
	c.name = name
	c.symbol = symbol
	c.decimals = decimals
	return c
}
