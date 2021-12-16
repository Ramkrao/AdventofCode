package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/Ramkrao/advent/utils"
)

func main() {
	lines := utils.ReadArrayFromFile("day16/input.txt")

	var bits []uint64
	for i := range lines {
		num, err := strconv.ParseUint(lines[i], 16, 32)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(utils.AsBits(num))
		bits = utils.AsBits(num)
	}

	version, _ := strconv.ParseInt(parseUint(bits[:3]), 2, 64)
	packet_type, _ := strconv.ParseInt(parseUint(bits[3:6]), 2, 64)

	// literal value
	if packet_type == 4 {
		fmt.Println(getLiteralValue(parseUint(bits[6:])))
	} else {

	}

	fmt.Println(version, packet_type)
}

func getLiteralValue(bits string) int64, err {
	var s bytes.Buffer
	for i := 0; i < len(bits); i++ {
		// fmt.Println(i, len(bits), string(bits[i]))
		// not a last group
		if bits[i] == '1' {
			s.WriteString(bits[i+1 : i+5])
		} else if bits[i] == '0' {
			s.WriteString(bits[i+1 : i+5])
			break
		}
		i = i + 4
	}
	// convert to decimal value
	d, _ := strconv.ParseInt(s.String(), 2, 64)
	return d, nil
}

func getOperatorValues(bits string) []int64 {
	values := make([]int64, 0)
	// check for length type
	if bits[0] == '0' {
		length, _  := strconv.ParseInt(bits[1:16], 2, 64)
		for i :=0; i < length; i++ {
			
		}
	}

	return values
}

func parseUint(bits []uint64) string {
	var s bytes.Buffer
	for _, b := range bits {
		s.WriteString(strconv.FormatUint(b, 10))
	}
	return s.String()
}
