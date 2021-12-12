package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func ReadArrayFromFile(filePath string) []string {

	fmt.Println("Start reading input file")

	// read the file content
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("Failed to read input", err)
	}
	// convert to string
	content := string(bytes)
	// parse the content to an array
	lines := strings.Split(content, "\n")

	fmt.Printf("Read %d lines from the file \n", len(lines))

	return lines
}

func StrArrToIntArr(arr []string) []int {
	out := make([]int, len(arr))
	for i, s := range arr {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		out[i] = val
	}
	return out
}

// Contains tells whether a contains x.
func Contains(arr []int, i int) bool {
	for _, n := range arr {
		if i == n {
			return true
		}
	}
	return false
}

// Contains tells whether arr contains s.
func ContainsStr(arr []string, s string) bool {
	for _, n := range arr {
		if s == n {
			return true
		}
	}
	return false
}

// Contains tells whether arr contains c.
func ContainsByte(arr []byte, c byte) bool {
	for _, n := range arr {
		if c == n {
			return true
		}
	}
	return false
}

func ComputeAdjacentPoints(arr [][]int, x int, y int) []string {
	// construct up, down, left, right and diagonals
	adj_points := make([]string, 0)
	// up
	if y > 0 {
		adj_points = append(adj_points, strconv.Itoa(x)+":"+strconv.Itoa(y-1))
	}
	// down
	if y < len(arr[0])-1 {
		adj_points = append(adj_points, strconv.Itoa(x)+":"+strconv.Itoa(y+1))
	}
	// left
	if x > 0 {
		adj_points = append(adj_points, strconv.Itoa(x-1)+":"+strconv.Itoa(y))
	}
	// right
	if x < len(arr)-1 {
		adj_points = append(adj_points, strconv.Itoa(x+1)+":"+strconv.Itoa(y))
	}
	// up-left
	if y > 0 && x > 0 {
		adj_points = append(adj_points, strconv.Itoa(x-1)+":"+strconv.Itoa(y-1))
	}
	// up-right
	if y > 0 && x < len(arr)-1 {
		adj_points = append(adj_points, strconv.Itoa(x+1)+":"+strconv.Itoa(y-1))
	}
	// down-left
	if y < len(arr[0])-1 && x > 0 {
		adj_points = append(adj_points, strconv.Itoa(x-1)+":"+strconv.Itoa(y+1))
	}
	// down-right
	if y < len(arr[0])-1 && x < len(arr)-1 {
		adj_points = append(adj_points, strconv.Itoa(x+1)+":"+strconv.Itoa(y+1))
	}
	// fmt.Println("Computed ", adj_points)
	return adj_points
}

// checks if the entire string is lower case
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func CountOccurance(arr []string, s string) int {
	count := 0
	for i := range arr {
		if s == arr[i] {
			count++
		}
	}
	return count
}
