package graph

import (
	"fmt"
	"testCode/structures/queue"
)

const (
	InitVertexNum = 10
)

type Graph struct {
	Node    []string
	Edge    [][]int
	NodeNum int
	EdgeNum int

	isDAG bool
	// 标记矩阵，0为当前顶点未访问，1为访问过，-1表示当前顶点访问过其他顶点。
	visited []int
}
//图的新建
func NewGraph() *Graph {
	//n * n 的邻接矩阵记录n 顶点集的连通关系
	edge := make([][]int,InitVertexNum)
	for i := 0; i< InitVertexNum;i++{
		edge[i] = make([]int,InitVertexNum)
	}

	return &Graph{
		Node:  make([]string, InitVertexNum),
		Edge:  edge,
		isDAG: true,
	}
}

//图的CURD
//1.新增顶点，存在返回false,成功返回true;自动扩容内部数据结构
func (g *Graph) AddNode(node string)  bool {
	if g.IndexOfNode(node) != -1 {
		return false //node exist
	}

	//超过则扩容
	if g.NodeNum >= cap(g.Node){
		g.ensureCapacity()
	}

	g.Node[g.NodeNum] = node
	g.NodeNum += 1
	return true
}

func (g *Graph) ensureCapacity() {
	oldCapacity := cap(g.Node)

	growVertex := make([]string,InitVertexNum)
	g.Node = append(g.Node,growVertex...) //go的append机制，在slice的后面添加默认值的slice

	//邻接矩阵edge: 新增行 new_n * count,原来的行slice 进行append
	newCapacity := cap(g.Node)
	growEdge := make([][]int,InitVertexNum)
	for i:=0;i<InitVertexNum;i++ {
		growEdge[i] = make([]int,newCapacity)
	}
	g.Edge = append(g.Edge,growEdge...)

	for i:=0;i<oldCapacity;i++ {
		edge := make([]int,InitVertexNum)
		g.Edge[i] = append(g.Edge[i],edge...)
	}
}


//2.新增边
func (g *Graph) AddEdge(fromVertex, toVertex string) bool {
	fromIndex := g.IndexOfNode(fromVertex)
	toIndex := g.IndexOfNode(toVertex)
	if fromIndex == -1 || toIndex == -1 {
		return false
	}

	g.Edge[fromIndex][toIndex] = 1
	if !g.isDAG {
		g.Edge[toIndex][fromIndex] = 1
	}
	g.EdgeNum++
	return true
}

//3.查找顶点，返回顶点在Graph 内部的数组下标
func (g *Graph) IndexOfNode(node string) int {
	for i:=0;i<cap(g.Node);i++{
		if g.Node[i] == node {
			return i
		}
	}

	return -1
}

func (g *Graph) GetNodebyIndex(index int) string{
	return g.Node[index]
}

//DFS
func (g *Graph) InitVisited() {
	g.visited = make([]int,cap(g.Node))
}

func (g *Graph) DFS(index int) {
	g.visited[index] = 1 //标记已经访问过
	for i:=0;i<g.NodeNum;i++{
		if g.Edge[index][i] != 0 && g.visited[i] != 1{
			g.DFS(i)
		}
	}
	fmt.Print("->"+g.Node[index])
}

//BFS
func (g *Graph) BFS(index int) {
	q := queue.NewNodeQueue()
	q.Enqueue(g.Node[index])

	for {
		if q.IsEmpty(){
			break
		}

		node := q.Dequeue()
		near := g.Edge[g.IndexOfNode(node)]
		fmt.Print("->"+node)
		for i:=0;i<len(near);i++{
			if near[i] == 1 && g.visited[i] != 1 {
				q.Enqueue(g.GetNodebyIndex(i))
				g.visited[i] = 1
			}
		}
	}
	fmt.Println()
}
