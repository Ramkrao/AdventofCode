package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

var alpha = regexp.MustCompile(`[a-z]{1}`)

var lines []string

func main() {
	lines = utils.ReadArrayFromFile("day24/input.txt")

	// part1
	fmt.Println(identifyLowestNumber("99999999999999"))

	// part2
	fmt.Println(identifyLowestNumber("11111111111111"))
}

func identifyLowestNumber(number string) (string, int) {
	lowestz := 999999999
	var low_num string
	for i := 0; i < 14; i++ {
		for j := 1; j < 10; j++ {
			var s string
			if i == 0 {
				s = strconv.Itoa(j) + number[i+1:]
			} else if i < 13 {
				s = number[:i] + strconv.Itoa(j) + number[i+1:]
			} else {
				s = number[:i] + strconv.Itoa(j)
			}
			ret := processInstr(s)
			if ret <= lowestz {
				lowestz = ret
				low_num = s
			}
		}
	}
	if lowestz != 0 {
		low_num, lowestz = identifyLowestNumber(low_num)
	}
	return low_num, lowestz
}

func processInstr(input string) int {
	var index, val int
	vars := make(map[string]int)
	for _, line := range lines {
		instr := strings.Split(line, " ")
		if len(instr) == 3 {
			val, _ = strconv.Atoi(instr[2])
			if len(alpha.FindString(instr[2])) > 0 {
				val = vars[instr[2]]
			}
		}
		switch instr[0] {
		case "inp":
			i, _ := strconv.Atoi(string(input[index]))
			vars[instr[1]] = i
			index++
		case "add":
			vars[instr[1]] = vars[instr[1]] + val
		case "mul":
			vars[instr[1]] = vars[instr[1]] * val
		case "div":
			if val != 0 {
				vars[instr[1]] = vars[instr[1]] / val
			}
		case "mod":
			if vars[instr[1]] >= 0 && val > 0 {
				vars[instr[1]] = vars[instr[1]] % val
			}
		case "eql":
			if vars[instr[1]] == val {
				vars[instr[1]] = 1
			} else {
				vars[instr[1]] = 0
			}
		}
	}
	return vars["z"]
}
