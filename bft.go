package main

import "fmt"


type Node struct {
	Name   string
	Status int     
	Votes  []*Node 

}

var nodes = make([]*Node, 0)

func createNodes() {
	A := Node{"A", 1, make([]*Node, 0)} 
	B := Node{"B", 1, make([]*Node, 0)} 
	C := Node{"C", 1, make([]*Node, 0)} 
	D := Node{"D", 0, make([]*Node, 0)} 
	nodes = append(nodes, &A)
	nodes = append(nodes, &B)
	nodes = append(nodes, &C)
	nodes = append(nodes, &D)

}

func votes() {
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		fmt.Println(node.Status)

		for j := 0; j < len(nodes); j++ {
			inode := nodes[j]
			node.Votes = append(node.Votes, inode)
		}

	}
}

func isValid() bool {
	node := nodes[len(nodes)-1]
	votes := node.Votes

	cnt := 0
	for _, n := range votes {
		fmt.Println(n.Status)
		if n.Status == 0 {
			cnt++
		}
	}

	if float32(cnt) < float32(len(nodes))/float32(3.0) {

		return true
	}

	return false

}

func main() {

	createNodes()
	votes()
	fmt.Println(isValid())
}
