package main

func main() {
	tree := merkleTree.MerkleTree{}
	printTree(tree.ContructionMerkleTree([]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e")})[0].(Noeud))
}
