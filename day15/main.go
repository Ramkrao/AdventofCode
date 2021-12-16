package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/Ramkrao/advent/utils"
	"github.com/RyanCarrier/dijkstra"
)

func main() {
	lines := utils.ReadArrayFromFile("day15/array.txt")
	f, _ := os.OpenFile("day15/input.txt", syscall.O_WRONLY, 0655)
	defer f.Close()

	cave := make([][]int, len(lines))
	for row := range cave {
		cave[row] = utils.StrArrToIntArr(strings.Split(lines[row], ""))
	}
	// expand the cave array
	cave = expandArray(cave, 5)
	fmt.Println(len(cave), len(cave[0]))

	// need a better logic, but for now something rudementry
	i := 0
	indexmap := make(map[string]int)
	for row := 0; row < len(cave); row++ {
		for col := 0; col < len(cave[0]); col++ {
			indexmap[strconv.Itoa(row)+":"+strconv.Itoa(col)] = i
			i++
		}
	}
	// fmt.Println(indexmap["499:499"])

	for row := range cave {
		for col := range cave[row] {
			// get adjacent points
			adjPoints := utils.ComputeAdjacentPoints(cave, row, col, true)
			f.WriteString(strconv.Itoa(indexmap[strconv.Itoa(row)+":"+strconv.Itoa(col)]))
			for _, point := range adjPoints {
				pos := utils.StrArrToIntArr(strings.Split(point, ":"))
				f.WriteString(" " + strconv.Itoa(indexmap[point]) + "," + strconv.Itoa(cave[pos[0]][pos[1]]))
			}
			f.WriteString("\n")
		}
	}
	fmt.Println("Finished writing to file")

	g, _ := dijkstra.Import("day15/input.txt")
	// fmt.Println(g)

	bestPath, _ := g.Shortest(0, len(cave)*len(cave)-1)
	fmt.Println(bestPath.Distance, bestPath.Path)
	// for i := range bestPaths {
	// 	fmt.Println(bestPaths[i].Distance, bestPaths[i].Path)
	// }
}

func expandArray(input [][]int, size int) [][]int {
	targetSize := len(input) * size
	// initialize target array
	target := make([][]int, len(input)*size)
	for row := range target {
		target[row] = make([]int, targetSize)
	}
	// assign initial value(s)
	for row := range input {
		for col := range input[row] {
			target[row][col] = input[row][col]
		}
	}
	// propogate values
	for row := range target {
		for col := range target[row] {
			// check the left first
			if col-len(input) >= 0 {
				target[row][col] = target[row][col-len(input)] + 1
				if target[row][col-len(input)]+1 > 9 {
					target[row][col] = 1
				}
			} else if row-len(input) >= 0 {
				target[row][col] = target[row-len(input)][col] + 1
				if target[row][col] > 9 {
					target[row][col] = 1
				}
			}
		}
	}
	return target
}
