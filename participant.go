package participant

type Participant struct {
	publicKey  byte
	privateKey byte
	ledger     float32
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(publicKey byte, privateKey byte, ledger float32) Participant {
	participant := Participant{publicKey, privateKey, ledger}
	return participant
}
