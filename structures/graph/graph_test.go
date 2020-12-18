package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")
	g.AddNode("F")

	g.AddEdge("A","B")
	g.AddEdge("A","C")
	g.AddEdge("A","D")
	g.AddEdge("B","E")
	g.AddEdge("C","E")
	g.AddEdge("E","F")

	g.InitVisited()
	g.DFS(0)
	fmt.Println("\n====================")
	g.InitVisited()
	g.BFS(0)
}