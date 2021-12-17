package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {
	dots := make(map[string][]int)

	lines := utils.ReadArrayFromFile("day13/input.txt")

	for l := range lines {
		if strings.Contains(lines[l], ",") {
			dots[lines[l]] = utils.StrArrToIntArr(strings.Split(lines[l], ","))
		} else if strings.Contains(lines[l], "fold") {
			fmt.Println(lines[l])
			vals := strings.Split(lines[l], "=")
			val, _ := strconv.Atoi(vals[1])
			dots = fold(dots, vals[0][len(vals[0])-1:], val)
			fmt.Println("Number of dots remaining ", len(dots))
		} else {
			fmt.Println("Number of dots to start ", len(dots))
		}
	}
	result := make([][]int, 6)
	for row := range result {
		result[row] = make([]int, 40)
	}
	for _, v := range dots {
		result[v[1]][v[0]] = 1
	}
	for row := range result {
		fmt.Println(result[row])
	}
}

func fold(dots map[string][]int, dir string, val int) map[string][]int {
	for k, v := range dots {
		if dir == "x" {
			if v[0] > val {
				s := strconv.Itoa(val-(v[0]-val)) + "," + strconv.Itoa(v[1])
				_, prs := dots[s]
				if !prs {
					dots[s] = []int{val - (v[0] - val), v[1]}
				}
				delete(dots, k)
			}
		} else if dir == "y" {
			if v[1] > val {
				s := strconv.Itoa(v[0]) + "," + strconv.Itoa(val-(v[1]-val))
				_, prs := dots[s]
				if !prs {
					dots[s] = []int{v[0], val - (v[1] - val)}
				}
				delete(dots, k)
			}
		}
	}
	return dots
}
