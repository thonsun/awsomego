package tree

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	nodeG := Node{item: "g", left: nil, right: nil}
	nodeF := Node{item: "f", left: &nodeG, right: nil}
	nodeE := Node{item: "e", left: nil, right: nil}
	nodeD := Node{item: "d", left: &nodeE, right: nil}
	nodeC := Node{item: "c", left: nil, right: nil}
	nodeB := Node{item: "b", left: &nodeD, right: &nodeF}
	nodeA := Node{item: "a", left: &nodeB, right: &nodeC}

	BreadFirstSearch(&nodeA)
	fmt.Println("\n===============")
	PreOrderSearch(&nodeA)
	fmt.Println("\n===============")
	MidOrderSearch(&nodeA)
	fmt.Println("\n===============")
	PostOrderSearch(&nodeA)
	fmt.Println("\n===============")
}
