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

type Node struct {
	left Hashable
	right Hashable
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

func (n Node) hash() Hash {
	var l, r [sha1.Size]byte
	l = n.left.hash()
	r = n.right.hash()
	return hash(append(l[:], r[:]...))
}

func buildMerkleTree(parts []Hashable) []Hashable { //[]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e")}
	var noeuds []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i+1 < len(parts) {
			noeuds = append(noeuds, Node {left: parts[i], right: parts[i+1]}) // par pairs on les ajoute au tableau noeuds
		}
	}
	if len(noeuds) == 1 {
		return noeuds
	} else { // ou sinon on refait
		return buildMerkleTree(noeuds)
	}
}

func printTree(node Node) {
	printNode(node, 0)
}

func printNode(node Node, level int) {
	fmt.Printf("(%d) %s %s\n", level, strings.Repeat(" ", level), node.hash())
	if l, ok := node.left.(Node); ok {
		printNode(l, level+1)
	} else if l, ok := node.left.(Block); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat(" ", level+1), l.hash(), l)
	}
	if r, ok := node.right.(Node); ok {
		printNode(r, level+1)
	} else if r, ok := node.right.(Block); ok {
		fmt.Printf("(%d) %s %s (data: %s)\n", level+1, strings.Repeat(" ", level+1), r.hash(), r)
	}
}

func main() {
	printTree(buildMerkleTree([]Hashable{Block("a"), Block("b"), Block("c"), Block("d"), Block("e"), Block("f")})[0].(Node))
}
