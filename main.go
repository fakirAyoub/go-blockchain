package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

// Type :
type Hash [20]byte
type Block string

type Noeud struct {
	gauche Hashable
	droite Hashable
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func hash(data []byte) Hash {
	return sha1.Sum(data)
}

type Hashable interface {
	hash() Hash
}

func (b Block) hash() Hash {
	return hash([]byte(b)[:])
}

func (n Noeud) hash() Hash {
	var l, r [sha1.Size]byte
	l = n.gauche.hash()
	r = n.droite.hash()
	return hash(append(l[:], r[:]...))
}

func ContructionMerkleTree(parts []Hashable) []Hashable { //[]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e")}
	var noeuds []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			noeuds = append(noeuds, Noeud{gauche: parts[i], droite: parts[i+1]}) // par pairs on les ajoute au tableau noeuds
		}
	}
	if len(noeuds) == 1 {
		return noeuds
	} else { // ou sinon on refait
		return ContructionMerkleTree(noeuds)
	}
}

func printTree(node Noeud) {
	printNode(node, 0)
}

func printNode(node Noeud, level int) {
	fmt.Printf("(%d) %s %s\n", level, strings.Repeat(" ", level), node.hash())
	if l, ok := node.gauche.(Noeud); ok {
		printNode(l, level+1)
	} else if l, ok := node.gauche.(Block); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat(" ", level+1), l.hash(), l)
	}
	if r, ok := node.droite.(Noeud); ok {
		printNode(r, level+1)
	} else if r, ok := node.droite.(Block); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat(" ", level+1), r.hash(), r)
	}
}

func main() {
	printTree(ContructionMerkleTree([]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e")})[0].(Noeud))
}
