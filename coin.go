package coin

type Coin struct {
	From     byte
	To       byte
	Value    int
	Data     byte
	GasPrice int
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(From byte, To byte, Value int, Data byte, GasPrice int) Coin {
	coin := Coin{From, To, Value, Data, GasPrice}
	return Coin
}
