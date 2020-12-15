package ledger

type Ledger struct {
	Transactions []Transaction
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(Transactions []Transaction) Ledger {
	ledger := Ledger{Transactions}
	return ledger
}
