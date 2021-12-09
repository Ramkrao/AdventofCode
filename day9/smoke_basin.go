package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

var path []string

func main() {
	lines := utils.ReadArrayFromFile("day9/input.txt")

	tubes := make([][]int, len(lines))
	for i := range tubes {
		tubes[i] = utils.StrArrToIntArr(strings.Split(lines[i], ""))
	}
	fmt.Println(tubes)

	result := 0
	basins := make([]int, 0)
	for i := range tubes {
		for j := range tubes[i] {
			if checkLowPoint(tubes, i, j) {
				fmt.Println("Found low point", i, j, tubes[i][j])
				result += tubes[i][j] + 1
				path = make([]string, 0)
				size := discoverBasin(tubes, i, j) + 1
				fmt.Println("Basin size ", size)
				basins = append(basins, size)
			}
		}
	}
	fmt.Println(result)
	sort.Ints(basins)
	fmt.Println(basins)
	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}

func checkLowPoint(tubes [][]int, x int, y int) bool {
	// construct up, down, left and right
	adj_vals := make([]int, 0)
	if y > 0 {
		adj_vals = append(adj_vals, tubes[x][y-1])
	}
	if y < len(tubes[0])-1 {
		adj_vals = append(adj_vals, tubes[x][y+1])
	}
	if x > 0 {
		adj_vals = append(adj_vals, tubes[x-1][y])
	}
	if x < len(tubes)-1 {
		adj_vals = append(adj_vals, tubes[x+1][y])
	}
	sort.Ints(adj_vals)
	return tubes[x][y] < adj_vals[0]
}

func computeAdjacentPoint(tubes [][]int, x int, y int) []string {
	// construct up, down, left and right
	adj_points := make([]string, 0)
	if y > 0 {
		adj_points = append(adj_points, strconv.Itoa(x)+":"+strconv.Itoa(y-1))
	}
	if y < len(tubes[0])-1 {
		adj_points = append(adj_points, strconv.Itoa(x)+":"+strconv.Itoa(y+1))
	}
	if x > 0 {
		adj_points = append(adj_points, strconv.Itoa(x-1)+":"+strconv.Itoa(y))
	}
	if x < len(tubes)-1 {
		adj_points = append(adj_points, strconv.Itoa(x+1)+":"+strconv.Itoa(y))
	}
	// fmt.Println("Computed ", adj_points)
	return adj_points
}

func discoverBasin(tubes [][]int, x int, y int) int {
	count := 0
	adj_points := computeAdjacentPoint(tubes, x, y)
	for _, point := range adj_points {
		xy := utils.StrArrToIntArr(strings.Split(point, ":"))
		if !utils.ContainsStr(path, strconv.Itoa(x)+":"+strconv.Itoa(y)) {
			path = append(path, strconv.Itoa(x)+":"+strconv.Itoa(y))
		}
		// fmt.Println("Path ", path)
		if tubes[x][y] < tubes[xy[0]][xy[1]] && tubes[xy[0]][xy[1]] != 9 {
			if !utils.ContainsStr(path, point) {
				fmt.Println("Found basin point ", xy[0], xy[1], tubes[xy[0]][xy[1]])
				count += discoverBasin(tubes, xy[0], xy[1])
				count++
			}
		}
	}
	return count
}
