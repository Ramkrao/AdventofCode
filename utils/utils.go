package utils

import (
	"fmt"
	"io/ioutil"
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
