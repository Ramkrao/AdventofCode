package main

import (
	"fmt"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	draws := "38,54,68,93,72,12,33,8,98,88,21,91,53,61,26,36,18,80,73,47,3,5,55,92,67,52,25,40,56,95,9,62,30,31,85,65,14,2,78,75,15,39,87,27,58,42,60,32,41,83,51,77,10,66,70,4,37,6,89,23,16,49,48,63,94,97,86,64,74,82,7,0,11,71,44,43,50,69,45,81,20,28,46,79,90,34,35,96,99,59,1,76,22,24,17,57,13,19,84,29"
	input := utils.StrArrToIntArr(strings.Split(draws, ","))

	totalBoards := 100
	boards := make([][][]int, totalBoards)
	for i := range boards {
		boards[i] = make([][]int, 5)
		for j := range boards[i] {
			boards[i][j] = make([]int, 5)
		}
	}

	lines := utils.ReadArrayFromFile("day4/input.txt")

	var x, y int
	for _, line := range lines {
		if line == "" {
			x++
			y = 0
			continue
		}
		boards[x][y] = utils.StrArrToIntArr(strings.Split(strings.TrimSpace(strings.Replace(line, "  ", " ", -1)), " "))
		y++
	}

	for i, draw := range input {
		fmt.Println("Checking draw", draw, "iteration", i)
		markBoards(boards, draw)
		results := checkBoards(boards)
		if len(results) > 0 {
			totalBoards = totalBoards - len(results)
			fmt.Println("Remaining boards", totalBoards)

			if totalBoards > 0 {
				for index := len(results) - 1; index >= 0; index-- {
					boards = remove(boards, results[index])
				}
			} else {
				var total int
				for j := range boards[0] {
					for k := range boards[0][j] {
						if boards[0][j][k] != -1 {
							total = total + boards[0][j][k]
						}
					}
				}
				fmt.Println("Final result", total, draw, total*draw)
				break
			}
		}
	}
}

func markBoards(boards [][][]int, draw int) {
	for i := range boards {
		for j := range boards[i] {
			for k := range boards[i][j] {
				if boards[i][j][k] == draw {
					boards[i][j][k] = -1
				}
			}
		}
	}
}

func checkBoards(boards [][][]int) []int {
	winningBoards := make([]int, 0)
	for i := range boards {
		for j := range boards[i] {
			if boards[i][j][0] == -1 && boards[i][j][1] == -1 && boards[i][j][2] == -1 && boards[i][j][3] == -1 && boards[i][j][4] == -1 {
				winningBoards = append(winningBoards, i)
				fmt.Println("Row win", j, i, boards[i])
			} else if boards[i][0][j] == -1 && boards[i][1][j] == -1 && boards[i][2][j] == -1 && boards[i][3][j] == -1 && boards[i][4][j] == -1 {
				winningBoards = append(winningBoards, i)
				fmt.Println("Column win", j, i, boards[i])
			}
		}
	}
	temp := make([]int, 0)
	for i := range winningBoards {
		if !contains(temp, winningBoards[i]) {
			temp = append(temp, winningBoards[i])
		}
	}
	if len(temp) > 0 {
		fmt.Println("Returning winning boards", temp)
	}
	return temp
}

func contains(arr []int, i int) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}
	return false
}

func remove(slice [][][]int, s int) [][][]int {
	return append(slice[:s], slice[s+1:]...)
}
