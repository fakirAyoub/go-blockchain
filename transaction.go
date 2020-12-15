package transaction

type Transaction struct {
	From     byte
	To       byte
	Value    int
	Data     byte
	GasPrice int
	Date     Date
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(From byte, To byte, Value int, Data byte, GasPrice int) Transaction {
	transaction := Transaction{From, To, Value, Data, GasPrice}
	return transaction
}
