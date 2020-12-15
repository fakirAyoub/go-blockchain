package block

type Block struct {
	Version      string
	Id           byte
	MerkelRoot   byte
	Date         Date
	Diff         float64
	Nonce        int
	MinedBy      *Miner
	Transactions []*Transaction
}

/*
Créer une instance de la classe Personnage

@return: struct Personnage
*/
func New(Version string, Id byte, MerkleRoot byte, Date Date, Diff float64, Nonce int, MinedBy *Miner, Transactions []*Transaction) Block {
	block := Block{Version, Id, MerkleRoot, Date, Diff, Nonce, MinedBy, Transactions}
	return block
}
