package main

type Coin struct {
	name string
	symbol rune
	decimal int32
	totalSupply int64
}

func (c *Coin) CoinConstructor (_name string, _symbol rune, _decimal int32) {
	c.name = _name;
	c.symbol = _symbol;
	c.decimal = _decimal;
}