package main

import "fmt"

// Reference - https://fodor.org/blog/go-graph/
type Caves struct {
	nodes []*CaveNode
}

type CaveNode struct {
	id    string
	edges []*CaveNode
}

// new instance of the graph
func New() *Caves {
	return &Caves{
		nodes: []*CaveNode{},
	}
}

// adds a node to the Cave struct
func (c *Caves) AddNode(id string) {
	// check and add only if it's not present
	for _, n := range c.nodes {
		if n.id == id {
			return
		}
	}
	c.nodes = append(c.nodes, &CaveNode{
		id:    id,
		edges: []*CaveNode{},
	})
}

// adds a directional edge between nodes
func (c *Caves) AddEdge(n1, n2 string) {
	var node1, node2 *CaveNode
	for _, n := range c.nodes {
		if n.id == n1 {
			node1 = n
		} else if n.id == n2 {
			node2 = n
		}
	}
	node1.edges = append(node1.edges, node2)
}

// get a list of connected nodes
func (c *Caves) Neighbours(id string) []*CaveNode {
	neighbours := []*CaveNode{}
	for _, n := range c.nodes {
		for _, edge := range n.edges {
			if n.id == id {
				neighbours = append(neighbours, n)
			}
			if edge.id == id {
				neighbours = append(neighbours, edge)
			}
		}
	}
	return neighbours
}

// list all the nodes and edges
func (c *Caves) NodesNEdges() {
	for _, node := range c.nodes {
		fmt.Println("Node ", node.id)
		for _, edge := range node.edges {
			fmt.Println(node.id, " -> ", edge.id)
		}
	}
}

// get a specific node by id
func (c *Caves) GetNodeByID(id string) *CaveNode {
	for _, node := range c.nodes {
		if node.id == id {
			return node
		}
	}
	return nil
}
