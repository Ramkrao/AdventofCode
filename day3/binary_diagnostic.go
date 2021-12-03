package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	zeroBits := make([]int, 12)
	oneBits := make([]int, 12)
	lines := utils.ReadArrayFromFile("day3/input.txt")

	for _, line := range lines {
		chars := strings.Split(line, "")
		for pos, char := range chars {
			if char == "0" {
				zeroBits[pos] = zeroBits[pos] + 1
			} else {
				oneBits[pos] = oneBits[pos] + 1
			}
		}
	}

	var prominentBits, nonPromBits strings.Builder
	for count, _ := range zeroBits {
		if zeroBits[count] > oneBits[count] {
			prominentBits.WriteString("0")
			nonPromBits.WriteString("1")
		} else {
			prominentBits.WriteString("1")
			nonPromBits.WriteString("0")
		}
	}

	fmt.Println(zeroBits, oneBits)
	fmt.Println(prominentBits.String(), nonPromBits.String())

	var gammaRate, epsilonRate int64
	if i, err := strconv.ParseInt(prominentBits.String(), 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
		gammaRate = i
	}
	if i, err := strconv.ParseInt(nonPromBits.String(), 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
		epsilonRate = i
	}
	fmt.Println("Power consumption", gammaRate*epsilonRate)

	var oxyRate, co2Rate int64
	oxy := computeRatings(lines, 0)
	if i, err := strconv.ParseInt(oxy, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
		oxyRate = i
	}

	co2 := computeRatings(lines, 1)
	if i, err := strconv.ParseInt(co2, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
		co2Rate = i
	}
	fmt.Println("Oxygen Generater Rating is", oxyRate*co2Rate)
}

func computeRatings(lines []string, gasType int) string {

	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)
	for count := 0; count < 12; count++ {

		bit := getPositionalBit(linesCopy, count, gasType)

		tempArray := make([]string, 0)
		for _, line := range linesCopy {
			if string(line[count]) == bit {
				tempArray = append(tempArray, line)
			}
		}
		linesCopy = tempArray
		fmt.Println(len(linesCopy), bit)
		if len(linesCopy) == 1 {
			break
		}
	}

	fmt.Println(linesCopy)
	return linesCopy[0]
}

func getPositionalBit(lines []string, pos int, gasType int) string {
	var zeroBit, oneBit int
	var char string
	for _, line := range lines {
		index := line[pos]
		if index == '0' {
			zeroBit = zeroBit + 1
		} else {
			oneBit = oneBit + 1
		}
	}
	if gasType == 0 {
		if oneBit >= zeroBit {
			char = "1"
		} else {
			char = "0"
		}
	} else {
		if zeroBit <= oneBit {
			char = "0"
		} else {
			char = "1"
		}
	}
	fmt.Println(zeroBit, oneBit, char)
	return char
}
