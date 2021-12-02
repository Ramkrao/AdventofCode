package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input []int

func init() {
	fmt.Println("Start reading input file")

	// read the file content
	bytes, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Errorf("Failed to read input", err)
	}
	// convert to string
	content := string(bytes)
	// parse the content to an array
	lines := strings.Split(content, "\n")
	// initialize empty array
	input = make([]int, len(lines))
	// convert []string to []int
	for i := 0; i < len(lines); i++ {
		j, err := strconv.Atoi(lines[i])
		if err == nil {
			input[i] = j
		}
	}

	fmt.Printf("Read %d integers from the file \n", len(input))
}

func main() {

	// Staring execution
	fmt.Println("Analysing sonar inputs")
	// Declaring output variables
	var increased, decreased, nochange int

	// simple for loop to execute computation
	for count := 0; count < len(input)-3; count++ {

		currBlock := input[count] + input[count+1] + input[count+2]
		nextBlock := input[count+1] + input[count+2] + input[count+3]

		if currBlock < nextBlock {
			increased++
			fmt.Println(currBlock, "increased ", count+1)
		} else if currBlock == nextBlock {
			nochange++
			fmt.Println(currBlock, "nochange ", count+1)
		} else {
			decreased++
			fmt.Println(currBlock, "decreased ", count+1)
		}
	}

	// ptint the results
	fmt.Println("Total increased ", increased)
	fmt.Println("Total nochange", nochange)
	fmt.Println("Total decreased", decreased)
}
