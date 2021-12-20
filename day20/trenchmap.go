package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {
	lines := utils.ReadArrayFromFile("day20/input.txt")

	var image = make([][]string, 100)
	var algo string

	for row, line := range lines {
		if row == 0 {
			algo = line
		} else {
			image[row-1] = strings.Split(line, "")
		}
	}
	image = expandArray(image, 3)
	count := 0
	for count < 50 {
		temp := make([][]string, len(image))
		for i := range temp {
			temp[i] = make([]string, len(image[0]))
			for j := range temp[i] {
				if count%2 == 0 {
					temp[i][j] = "#"
				} else {
					temp[i][j] = "."
				}
			}
		}
		for row := 50; row < 250; row++ {
			for col := 50; col < 250; col++ {
				temp[row][col] = string(algo[getAdjPointsBinary(image, row, col)])
			}
		}
		image = temp
		count++
		pixels := 0
		fmt.Println("---------------------")
		for i := range image {
			for j := range image[i] {
				if image[i][j] == string('#') {
					pixels++
				}
			}
		}
		fmt.Println("Total light pixels ", pixels, count)
	}

}

func getAdjPointsBinary(image [][]string, x int, y int) int {
	// construct up, down, left, right and diagonals
	var binary bytes.Buffer
	// up-left
	if y > 0 && x > 0 {
		binary.WriteString(image[x-1][y-1])
	}
	// up
	if x > 0 {
		binary.WriteString(image[x-1][y])
	}
	// up-right
	if y < len(image[0])-1 && x > 0 {
		binary.WriteString(image[x-1][y+1])
	}
	// left
	if y > 0 {
		binary.WriteString(image[x][y-1])
	}
	// centre
	if x > 0 {
		binary.WriteString(image[x][y])
	}
	// right
	if y < len(image[0])-1 {
		binary.WriteString(image[x][y+1])
	}
	// down-left
	if y > 0 && x < len(image)-1 {
		binary.WriteString(image[x+1][y-1])
	}
	// down
	if x < len(image)-1 {
		binary.WriteString(image[x+1][y])
	}
	// down-right
	if y < len(image[0])-1 && x < len(image)-1 {
		binary.WriteString(image[x+1][y+1])
	}
	temp := binary.String()
	temp = strings.ReplaceAll(temp, ".", "0")
	temp = strings.ReplaceAll(temp, "#", "1")

	d, _ := strconv.ParseInt(temp, 2, 64)
	return int(d)
}

func expandArray(input [][]string, size int) [][]string {
	targetSize := len(input) * size
	// initialize target array
	target := make([][]string, len(input)*size)
	for row := range target {
		target[row] = make([]string, targetSize)
	}
	// assign initial value(s)
	for row := range input {
		for col := range input[row] {
			target[row+len(input)][col+len(input)] = input[row][col]
		}
	}
	for i := range target {
		for j := range target[i] {
			if target[i][j] != "#" {
				target[i][j] = "."
			}
		}
	}
	return target
}
