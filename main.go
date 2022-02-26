package main

import (
    "fmt"
)
type Graph struct {
    nodes []*Node
}

type Node struct {
	key int
	neighbors []*Node
}

func (g *Graph) AddNode(k int) {
    g.nodes = append(g.nodes, &Node{key: k})
}

func (g *Graph) AddEdge(from, to int) {
    fromNode := g.getNode(from)
    toNode := g.getNode(to)
    fromNode.neighbors = append(fromNode.neighbors, toNode)
}


func (g *Graph) getNode(k int) *Node{
	for _, v := range g.nodes {
		if v.key == k {
			return v
		}
	}
	return nil
}

/*
Use Topological Sort to get MaxDistance from Root
*/
func (g *Graph) getMaxDistance(root *Node) int {
	inEdges := make(map[*Node]int)
	queue := []*Node{}
	distance := make(map[*Node]int)
	finalMax := 0

	for _, v := range g.nodes {
		for _, neighbor := range v.neighbors {
			inEdges[neighbor]++
		}
	}

    queue = append(queue, root)
   
    for len(queue) > 0 {
        v := queue[0]
        queue = queue[1:]

        for _, neighbor := range v.neighbors {
        	inEdges[neighbor]--
        	if inEdges[neighbor] == 0 {
        		queue = append(queue, neighbor)
                distance[neighbor] = max(distance[neighbor], distance[v]+1)
                finalMax = max(finalMax, distance[neighbor])
        	}
        }
    }
    return finalMax
 }
	

func max(i, j int) int {
    if i > j {
    	return i
    }
    return j
}

func (g *Graph) Print() {
	for _, node := range g.nodes {
		fmt.Printf("\n Node %v: ", node.key)
		for _, v := range node.neighbors {
			fmt.Printf(" %v: ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	g := Graph{}
    g.AddNode(1)
    g.AddNode(2)
    g.AddNode(3)
    g.AddNode(4)
    g.AddNode(5)


    g.AddEdge(1, 2)
    g.AddEdge(1, 3)
    g.AddEdge(1, 4)
    g.AddEdge(2, 4)
    g.AddEdge(3, 4)
    g.AddEdge(4, 5)
    g.AddEdge(3, 5)

    g.Print()

    root := g.getNode(1)

    fmt.Println("Max Distance From Root:", g.getMaxDistance(root))
}

