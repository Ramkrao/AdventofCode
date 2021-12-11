package main

import (
	"fmt"
	"sort"

	"github.com/Ramkrao/advent/utils"
)

var startingChar = []byte{40, 91, 123, 60}
var endingChar = []byte{41, 93, 125, 62}

func main() {
	scores := make([]int, 0)
	lines := utils.ReadArrayFromFile("day10/input.txt")
	for i := range lines {
		fmt.Println(lines[i])
		if checkForCorruption([]byte(lines[i])) == 0 {
			missing := fixIncompleteLines([]byte(lines[i]))
			scores = append(scores, computePoints(missing))
		}
	}
	sort.Ints(scores)
	fmt.Println(scores, len(scores), len(scores)/2)
	fmt.Println("Middle score : ", scores[len(scores)/2])
}

func computePoints(bytes []byte) int {
	score := 0
	for _, c := range bytes {
		score *= 5
		switch c {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}
	return score
}

func fixIncompleteLines(bytes []byte) []byte {
	chunks := make([]byte, 0)
	missing := make([]byte, 0)

	for _, c := range bytes {
		if utils.ContainsByte(startingChar, c) {
			chunks = append(chunks, c)
		} else if utils.ContainsByte(endingChar, c) {
			lastChar := chunks[len(chunks)-1]
			if c-lastChar <= 2 {
				chunks = chunks[:len(chunks)-1]
			}
		}
	}
	for i := len(chunks) - 1; i >= 0; i-- {
		var m byte
		switch chunks[i] {
		case '(':
			m = ')'
		case '[':
			m = ']'
		case '{':
			m = '}'
		case '<':
			m = '>'
		}
		chunks = append(chunks, m)
		missing = append(missing, m)
	}
	fmt.Println("Is it complete ", checkForCorruption(chunks) == 0)
	return missing
}

func checkForCorruption(bytes []byte) byte {
	chunks := make([]byte, 0)
	for _, c := range bytes {
		if utils.ContainsByte(startingChar, c) {
			chunks = append(chunks, c)
		} else if utils.ContainsByte(endingChar, c) {
			lastChar := chunks[len(chunks)-1]
			if c-lastChar > 2 {
				return c
			} else {
				chunks = chunks[:len(chunks)-1]
			}
		}
	}
	return 0
}
