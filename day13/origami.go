package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	lines := utils.ReadArrayFromFile("day13/input.txt")

	paper := make([][]string, 12) //15) //894)
	for row := range paper {
		paper[row] = make([]string, 16) //11) //1311)
	}

	for i := range lines {
		if strings.Contains(lines[i], ",") {
			points := utils.StrArrToIntArr(strings.Split(lines[i], ","))
			paper[points[1]][points[0]] = "#"
		} else if strings.Contains(lines[i], "fold") {
			fold := strings.Split(lines[i], "=")
			val, _ := strconv.Atoi(fold[1])
			paper = foldPaper(paper, string(fold[0][11]), val)
			fmt.Println("----------------step1---------------", i)
			count := 0
			for row := range paper {
				for col := range paper[row] {
					if paper[row][col] == "#" {
						count++
					}
				}
				fmt.Println(paper[row])
			}
			fmt.Println("Count ", count)
		} else {
			fmt.Println("----------------step2---------------", i)
			for row := range paper {
				for col := range paper[row] {
					if paper[row][col] != "#" {
						paper[row][col] = "."
					}
				}
				fmt.Println(paper[row])
			}
		}
	}
	// for row := range paper {
	// 	fmt.Println(paper[row])
	// }
}

func foldPaper(paper [][]string, direction string, val int) [][]string {
	var temp [][]string
	fmt.Println(direction, "=", val)
	if direction == "x" {
		temp = make([][]string, len(paper))
		var x, y, z int
		for row := range temp {
			k := len(paper[0]) - 1
			temp[row] = make([]string, val)
			for col := range temp[row] {

				if paper[row][col] == "#" {
					x++
				}
				if paper[row][k] == "#" {
					y++
				}
				if paper[row][col] == "#" && paper[row][k] == "#" {
					z++
				}
				if paper[row][col] == "#" || paper[row][k] == "#" {
					temp[row][col] = "#"
				} else {
					temp[row][col] = "."
				}
				if temp[row][col] == "#" {
					fmt.Println(row, col, k, paper[row][col], paper[row][k])
				}
				k--
			}
		}
		fmt.Println(x, y, z)
	} else if direction == "y" {
		temp = make([][]string, val)
		i := len(paper) - 1
		var x, y, z int
		for row := range temp {
			temp[row] = make([]string, len(paper[0]))
			for col := range temp[row] {

				if paper[row][col] == "#" {
					x++
				}
				if paper[i][col] == "#" {
					y++
				}
				if paper[row][col] == "#" && paper[i][col] == "#" {
					z++
				}
				if paper[row][col] == "#" || paper[i][col] == "#" {
					temp[row][col] = "#"
				} else {
					temp[row][col] = "."
				}
				if temp[row][col] == "#" {
					fmt.Println(row, col, i, paper[row][col], paper[i][col])
				}
			}
			i--
		}
		fmt.Println(x, y, z)
	}
	return temp
}
