package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	var x, y, z int
	lines := utils.ReadArrayFromFile("day2/input.txt")

	for _, line := range lines {

		instr := strings.Split(line, " ")
		val, _ := strconv.Atoi(instr[1])
		switch instr[0] {
		case "forward":
			x = x + val
			y = y + (z * val)
		case "down":
			z = z + val
		case "up":
			z = z - val
		}
	}

	fmt.Println(x, y, x*y)
}
