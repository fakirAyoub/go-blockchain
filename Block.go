package main

import (
	"bytes"
	"crypto/sha256"
)

//Block
type Block struct {
	//hash du block
	Hash []byte
	//date de création
	Timestamp int64
	//données du block
	Transactions []*Transaction
	//hash du block précédent
	PrevHash []byte
	Nonce int32
}

//Blockchain avec un tableau contenant les pointeurs vers chaque block
type Blockchain struct {
	Blocks []*Block
}
//création d'un hash
func (b *Block) HashingBlock() {
	//join prend un tableau 2d avec le hash du block et ses données puis les convertit
	blockData := bytes.Join([][]byte{b.Hash, b.Timestamp, b.Transactions, b.PrevHash}, b.Nonce)
	// fmt.Println(blockData)
	blockHash := sha256.Sum256(blockData)
	// fmt.Println(blockHash)
	b.Hash = blockHash[:]
}

//création d'un block
//données du block, hash du block précédent, retourne le pointeur du block créé
func CreateBlock(data string, previousHash []byte, nonce int32, transactions []*Transaction) *Block {
	//créé un block avec sa référence
	//conversion des données en bytes
	block := &Block{[]byte{}, now.Unix() , transaction, previousHash, nonce}
	//hash le block
	block.HashingBlock()
	return block
}

//ajout d'un block
func (chain *Blockchain) AddBlock(data string) {
	//récupération du block précédent
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	//création du nouveau block
	newBlock := CreateBlock(data, previousBlock.Hash)
	//ajout du block à la blockchain
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockchain() *Blockchain {
	//retourne l'adresse de la Blockchain avec un tableau de bloc et le block genesis
	return &Blockchain{[]*Block{Genesis()}}
}