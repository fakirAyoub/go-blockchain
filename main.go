package main

func main() {
	printTree(ContructionMerkleTree([]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e")})[0].(Noeud))
}
