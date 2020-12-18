package tree

import (
	"fmt"
)

type Node struct {
	item string
	left *Node
	right *Node
}


func BreadFirstSearch(node *Node) {
	q := []*Node{node}
	for len(q)>0 {
		node := q[0]
		q = q[1:]
		fmt.Print(node.item)
		if node.left != nil {
			q = append(q,node.left)
		}
		if node.right != nil{
			q = append(q,node.right)
		}
	}
}

func DepthFirstSearch(node *Node){
	if node.left != nil {
		DepthFirstSearch(node.left)
	}
	fmt.Print(node.item)
	if node.right != nil {
		DepthFirstSearch(node.right)
	}
}

func PreOrderSearch(node *Node){
	fmt.Print(node.item)
	if node.left != nil {
		PreOrderSearch(node.left)
	}
	if node.right != nil {
		PreOrderSearch(node.right)
	}
}

func MidOrderSearch(node *Node){
	if node.left != nil {
		PreOrderSearch(node.left)
	}
	fmt.Print(node.item)
	if node.right != nil {
		PreOrderSearch(node.right)
	}
}

func PostOrderSearch(node *Node){
	if node.left != nil {
		PreOrderSearch(node.left)
	}
	if node.right != nil {
		PreOrderSearch(node.right)
	}
	fmt.Print(node.item)
}


