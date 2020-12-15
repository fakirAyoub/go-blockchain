package miner

type Miner struct {
	Addr        byte
	BlocksMined int
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(Addr byte, BlocksMined int) Miner {
	miner := Miner{Addr, BlocksMined}
	return miner
}
