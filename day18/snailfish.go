package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {
	lines := utils.ReadArrayFromFile("day18/input.txt")

	var join string
	results := map[int]string{}
	for row := 0; row < len(lines); row++ {
		for j := 1; j < len(lines); j++ {
			join = fmt.Sprintf("[%s,%s]", lines[row], lines[j])
			// process until done
			isprocessed, str := process(join)
			join = str
			for !isprocessed {
				isprocessed, join = process(join)
			}
			results[computeMagnitude(join)] = join

			// results could vary when joined the other way, so doing both
			join = fmt.Sprintf("[%s,%s]", lines[j], lines[row])
			// process until done
			isprocessed, str = process(join)
			join = str
			for !isprocessed {
				isprocessed, join = process(join)
			}
			results[computeMagnitude(join)] = join
		}
	}
	var magnitudes []int
	for k, _ := range results {
		magnitudes = append(magnitudes, k)
	}
	sort.Ints(magnitudes)
	fmt.Println("Largest magnitude : ", magnitudes[len(magnitudes)-1], len(magnitudes))
}

func process(join string) (bool, string) {
	explode_index := getFirstExplodingPair(join)
	for explode_index > 0 {
		join = explode(join, explode_index, strings.Index(join[explode_index:], "]")+explode_index+1)
		explode_index = getFirstExplodingPair(join)
	}
	if doubledigitnumeric.FindAllStringIndex(join, -1) != nil {
		join = checkAndSplit(join)
	}
	explode_index = getFirstExplodingPair(join)
	if explode_index < 0 && doubledigitnumeric.FindAllStringIndex(join, -1) == nil {
		return true, join
	}
	return false, join
}

func getFirstExplodingPair(bytes string) int {
	chunks := make([]rune, 0)
	for i, c := range bytes {
		if '[' == c {
			chunks = append(chunks, c)
			if len(chunks) == 5 {
				// just make sure to pick the whole numbers array
				index := wholearray.FindStringIndex(bytes[i:])
				if len(index) > 0 {
					// fmt.Println("Found the leftmost exploding pair", i+index[0])
					return i + index[0]
				}
			}
		} else if ']' == c {
			chunks = chunks[:len(chunks)-1]
		}
	}
	return -1
}

var wholearray = regexp.MustCompile(`\[[0-9]+,[0-9]+\]`)
var numeric = regexp.MustCompile(`[0-9]+`)
var doubledigitnumeric = regexp.MustCompile(`[0-9]{2}`)

func explode(line string, start, end int) string {
	// fmt.Println("explode array", line, start, end, line[start+1:end-1])
	// store the numbers to be exploded
	numbers := utils.StrArrToIntArr(strings.Split(line[start+1:end-1], ","))
	// replace the exploding block with 0
	line = line[:start] + "0" + line[end:]
	// next find the left and right numbers to add
	indexes := numeric.FindAllStringIndex(line, -1)
	var leftnumindex, rightnumindex []int
	for i := range indexes {
		if indexes[i][0] < start {
			leftnumindex = indexes[i]
		}
	}
	for i := len(indexes) - 1; i > 0; i-- {
		if indexes[i][0] > start {
			rightnumindex = indexes[i]
		}
	}
	// increment the values, if any
	if len(leftnumindex) > 0 {
		leftnum, err := strconv.Atoi(string(line[leftnumindex[0]:leftnumindex[1]]))
		if err != nil {
			fmt.Println("leftnum - error parsing ", string(line[leftnumindex[0]:leftnumindex[1]]), err)
			os.Exit(-1)
		}
		line = line[:leftnumindex[0]] + strconv.Itoa(leftnum+numbers[0]) + line[leftnumindex[1]:]
		len_newleftnum := len(strconv.Itoa(leftnum + numbers[0]))
		if len_newleftnum > len(strconv.Itoa(leftnum)) && len(rightnumindex) > 0 {
			rightnumindex[0] += len_newleftnum - len(strconv.Itoa(leftnum))
			rightnumindex[1] += len_newleftnum - len(strconv.Itoa(leftnum))
		}
	}
	if len(rightnumindex) > 0 {
		rightnum, err := strconv.Atoi(string(line[rightnumindex[0]:rightnumindex[1]]))
		if err != nil {
			fmt.Println("rightnum - error parsing ", string(line[rightnumindex[0]:rightnumindex[1]]), err)
			os.Exit(-1)
		}
		line = line[:rightnumindex[0]] + strconv.Itoa(rightnum+numbers[1]) + line[rightnumindex[1]:]
	}
	return line
}

func checkAndSplit(line string) string {
	// fmt.Println("Check for split ", line)
	// find if there are any >9 digits
	indexes := doubledigitnumeric.FindAllStringIndex(line, -1)
	if len(indexes) > 0 {
		val, err := strconv.Atoi(line[indexes[0][0]:indexes[0][1]])
		if err != nil {
			fmt.Println("checkAndSplit - error parsing ", line[indexes[0][0]:indexes[0][1]], err)
			os.Exit(-1)
		}
		line = fmt.Sprintf("%s[%d,%d]%s", string(line[:indexes[0][0]]), val/2, val/2+val%2, string(line[indexes[0][1]:]))
		// fmt.Println("After split :: ", line)
	}
	return line
}

func computeMagnitude(line string) int {
	for wholearray.FindStringIndex(line) != nil {
		index := wholearray.FindStringIndex(line)
		vals := utils.StrArrToIntArr(strings.Split(line[index[0]+1:index[1]-1], ","))
		line = line[:index[0]] + strconv.Itoa(vals[0]*3+vals[1]*2) + line[index[1]:]
	}
	magnitude, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return magnitude
}
