package main

import (
	"fmt"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

var paths []string
var endNodes = []string{"start", "end"}

func main() {
	// read from the file
	lines := utils.ReadArrayFromFile("day12/input.txt")

	// instantiate the caves graph
	caves := New()
	// process the input
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		// Add the nodes to the graph
		caves.AddNode(nodes[0])
		caves.AddNode(nodes[1])
		// Create an edge between nodes
		// For start and end node, it's one direction
		if nodes[0] == "start" || nodes[1] == "end" {
			caves.AddEdge(nodes[0], nodes[1])
		} else if nodes[1] == "start" || nodes[0] == "end" {
			caves.AddEdge(nodes[1], nodes[0])
		} else {
			caves.AddEdge(nodes[0], nodes[1])
			caves.AddEdge(nodes[1], nodes[0])
		}
	}
	// caves.NodesNEdges()
	//get the start node by id
	start := caves.GetNodeByID("start")
	for _, n := range start.edges {
		paths := findPaths(caves, n, []string{"start"})
		fmt.Println("Found total path ", len(paths))
	}
}

func findPaths(caves *Caves, node *CaveNode, path []string) []string {
	path = append(path, node.id)
	for _, n := range node.edges {
		// fmt.Println("Traversing", n.id, path, checkSmallCaveCondition(path, n.id))
		// If the current node is not the end, and not a lower cave and been visited already
		if n.id == "end" { // stop when reaching end node
			paths = append(paths, strings.Join(append(path, n.id), ","))
		} else if checkSmallCaveCondition(path, n.id) {
			findPaths(caves, n, path)
		}
	}
	return paths
}

func checkSmallCaveCondition(path []string, id string) bool {
	// first it shouldn't be end
	if id != "end" {
		// next check if it's lower cave
		if utils.IsLower(id) {
			//if not visited, return true
			if !utils.ContainsStr(path, id) {
				return true
			}
			// finally check only one lower cave is visited twice
			lowerCaves := getLowerCaseStrings(path)
			for i := range lowerCaves {
				if utils.ContainsStr(path, lowerCaves[i]) &&
					utils.CountOccurance(path, lowerCaves[i]) > 1 {
					return false
				}
			}
		}
	}
	return true
}

func getLowerCaseStrings(arr []string) []string {
	temp := make([]string, 0)
	for i := range arr {
		if utils.IsLower(arr[i]) &&
			!utils.ContainsStr(temp, arr[i]) &&
			!utils.ContainsStr(endNodes, arr[i]) {
			temp = append(temp, arr[i])
		}
	}
	return temp
}
