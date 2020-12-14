package main

import (
	"bytes"
	"crypto/sha256"
)

//transaction
type Transaction struct {
	//hash du sender
	sender []byte
	//hash du receiver
	receiver []byte
	//montant de la transaction
	amount int
}

//Block
type Block struct {
	//hash du block
	Hash []byte
	//données du block
	Transactions []*Transaction
	//hash du block précédent
	PrevHash []byte
	Nonce int
}

//MerkelTree
type MerkelTree struct {
	//noeud root
	NodeRoot *MerkelNode
}

type MerkelNode struct {
	//noeud gauche
	Right *MerkelNode
	//noeud droit
	Left  *MerkelNode
	//données du noeud
	Data  []byte
}


//création d'un noeud avec noeud de droite/gauche et les données en paramètre, renvoue un noeud
func CreateNode(right *MerkelNode, left *MerkelNode, data []byte) *MerkelNode{
	node := MerkelNode{}

	if right == nil && left == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else{
		hash := sha256.Sum256(append(left.Data, right.Data ...))
		node.Data = hash[:]
	}
	node.Right = right
	node.Left = left

	return &node
}



func CreateMerkelTree(data []byte) *MerkelTree {
	var nodes []MerkelNode

	if len(data)%2 != 0{
		return nil
	}

	for _, d := range data{
		node := CreateNode(nil, nil, d)
		nodes = append(nodes, *node)
	}
	
	layer := []MerkelNode
	do {
		for i := 0; i <len(nodes); i +=2{
			node = CreateNode(&nodes[i], &nodes[i+1], nil)
			layer = append(layer, *node)
		}
		nodes = layer
	}while(len(layer) != 1)

	tree := MerkelTree{&nodes}
	return &tree
}



//création d'un hash
//func (b *Block) HashingBlock() {
//	//join prend un tableau 2d avec le hash du block et ses données puis les convertit
//	blockData := bytes.Join([][]byte{b.Hash, b.Transactions}, []byte{})
//	// fmt.Println(blockData)
//	blockHash := sha256.Sum256(blockData)
//	// fmt.Println(blockHash)
//	b.Hash = blockHash[:]
//}
////création d'un block
////données du block, hash du block précédent, retourne le pointeur du block créé
//func CreateBlock(Transactions []Transaction, previousHash []byte) *Block {
//	//créé un block avec sa référence
//	//conversion des données en bytes
//	block := &Block{[]byte{}, []byte(data), previousHash}
//	//hash le block
//	block.HashingBlock()
//	return block
//}
