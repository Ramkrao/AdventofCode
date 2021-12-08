package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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
