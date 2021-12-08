package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

var segments []int = []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func main() {
	lines := utils.ReadArrayFromFile("day8/input.txt")

	num_uniq_digits := 0
	for i := range lines {
		num_uniq_digits += countUniqueInstances(strings.Split(strings.TrimSpace(strings.Split(lines[i], "|")[1]), " "))
	}
	fmt.Println("Unique digits", num_uniq_digits)
	var output string

	// TODO: very poor coding, needs refactoring
	for _, line := range lines {
		output = ""
		misaligned := false
		sequence := constructSequence(strings.Split(strings.TrimSpace(strings.Split(line, "|")[0]), " "))
		digits := strings.Split(strings.TrimSpace(strings.Split(line, "|")[1]), " ")
		for _, digit := range digits {
			val := recognizeNumber(sequence, digit)
			if val == -1 {
				misaligned = true
			}
			output = output + strconv.Itoa(val)
		}
		// swap positions 2 and 5
		if misaligned {
			output = ""
			misaligned = false
			temp := sequence[2]
			sequence[2] = sequence[5]
			sequence[5] = temp
			for _, digit := range digits {
				val := recognizeNumber(sequence, digit)
				if val == -1 {
					misaligned = true
				}
				output = output + strconv.Itoa(val)
			}
		}
		// swap positions 1 and 3, 2 and 5
		if misaligned {
			output = ""
			misaligned = false
			temp := sequence[1]
			sequence[1] = sequence[3]
			sequence[3] = temp
			for _, digit := range digits {
				val := recognizeNumber(sequence, digit)
				if val == -1 {
					misaligned = true
				}
				output = output + strconv.Itoa(val)
			}
		}
		// swap positions 1 and 3
		if misaligned {
			output = ""
			misaligned = false
			temp := sequence[2]
			sequence[2] = sequence[5]
			sequence[5] = temp
			for _, digit := range digits {
				val := recognizeNumber(sequence, digit)
				if val == -1 {
					misaligned = true
				}
				output = output + strconv.Itoa(val)
			}
		}
		if !misaligned {
			fmt.Println(output)
		}
	}

}

func countUniqueInstances(digits []string) int {
	count := 0
	unique_list := []int{2, 4, 3, 7}
	fmt.Println("Processing ", digits)
	for i := range digits {
		if utils.Contains(unique_list, len(digits[i])) {
			count++
		}
	}
	return count
}

func constructSequence(sequence []string) []string {

	display_arr := make([]string, 7)
	// sort the input array based on string length
	sort.Slice(sequence, func(i, j int) bool {
		return len(sequence[i]) < len(sequence[j])
	})

	// process number 1
	if len(sequence[0]) == 2 {
		display_arr[2] = string(sequence[0][0])
		display_arr[5] = string(sequence[0][1])
	}
	// process number 7
	if len(sequence[1]) == 3 {
		display_arr[0] = getOddChar(sequence[1], sequence[0])[0]
	}
	// process number 4
	if len(sequence[2]) == 4 {
		display_arr[1] = getOddChar(sequence[2], sequence[0])[0]
		display_arr[3] = getOddChar(sequence[2], sequence[0])[1]
	}
	// process 2,3,5
	common := getCommonChar(sequence[3], sequence[4], sequence[5])
	display_arr[6] = getOddChar(strings.Join(common, ""), strings.Join(display_arr, ""))[0]
	// Find the last missing character
	display_arr[4] = getOddChar("abcdefg", strings.Join(display_arr, ""))[0]

	return display_arr
}

func getCommonChar(a string, b string, c string) []string {
	common := make([]string, 0)
	for _, ch := range strings.Split(a, "") {
		if strings.Contains(b, ch) && strings.Contains(c, ch) {
			common = append(common, ch)
		}
	}
	return common
}

func getOddChar(a string, b string) []string {
	odd := make([]string, 0)
	for _, c := range strings.Split(a, "") {
		if !strings.Contains(b, c) {
			odd = append(odd, c)
		}
	}
	return odd
}

// TODO: very poor coding, needs refactoring
func recognizeNumber(chars []string, input string) int {
	// recognize 1
	if len(input) == 2 &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[5]) {
		return 1
	} else if len(input) == 3 && // recognize 7
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[5]) {
		return 7
	} else if len(input) == 4 && // recognize 4
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[5]) {
		return 4
	} else if len(input) == 5 && //recognize 2
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[4]) &&
		strings.Contains(input, chars[6]) {
		return 2
	} else if len(input) == 5 && //recognize 3
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 3
	} else if len(input) == 5 && //recognize 5
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 5
	} else if len(input) == 6 && //recognize 0
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[4]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 0
	} else if len(input) == 6 && //recognize 6
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[4]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 6
	} else if len(input) == 6 && //recognize 9
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 9
	} else if len(input) == 7 && //recognize 8
		strings.Contains(input, chars[0]) &&
		strings.Contains(input, chars[1]) &&
		strings.Contains(input, chars[2]) &&
		strings.Contains(input, chars[3]) &&
		strings.Contains(input, chars[4]) &&
		strings.Contains(input, chars[5]) &&
		strings.Contains(input, chars[6]) {
		return 8
	}
	return -1
}
