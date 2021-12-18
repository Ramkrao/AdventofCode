package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/Ramkrao/advent/utils"
)

var hexToBinaryDict = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type packet struct {
	version     int64
	packet_type int64
	length_type byte
	value       int64
	subpackets  []packet
	length      int
	orig_str    string
}

func (p *packet) toString() {
	fmt.Fprintf(os.Stdout, "Packet info :: %#v\n", p)
}

func main() {
	lines := utils.ReadArrayFromFile("day16/input.txt")
	for i := range lines {
		binary := hexStrToBinaryStr(lines[i])
		fmt.Println("------------------------------------------------------------")
		fmt.Println("Processing line ", binary)
		p := packet{}
		p.orig_str = binary
		packet := parsePacket(binary, &p, false)
		fmt.Println("Final packet details ::")
		packet.toString()
	}
}

func parsePacket(bits string, pkt *packet, isSubPacket bool) *packet {
	fmt.Println("********Parsing BITS************", bits)
	p := pkt
	if isSubPacket {
		p = &packet{}
	}
	for count := 0; count < len(bits); count++ {
		startindex := count
		fmt.Println("Remaining BITS ", bits[count:], count)
		// this is to consider the padding bits
		if count+5 >= len(bits) {
			return pkt
		}
		version, _ := strconv.ParseInt(bits[count:count+3], 2, 64)
		p.version = version
		val, _ := strconv.ParseInt(bits[count+3:count+6], 2, 64)
		p.packet_type = val
		// packet.toString()
		// literal value
		if p.packet_type == 4 {
			endindex := 0
			var s bytes.Buffer
			for i := count + 6; i < len(bits); i++ {
				fmt.Println(i, string(bits[i]))
				// not a last group
				if bits[i] == '1' {
					s.WriteString(bits[i+1 : i+5])
				} else if bits[i] == '0' {
					s.WriteString(bits[i+1 : i+5])
					endindex = i + 4
					break
				}
				i = i + 4
			}
			count = endindex
			fmt.Println("count :", count)
			// convert to decimal value
			d, _ := strconv.ParseInt(s.String(), 2, 64)
			p.value = d
			p.length = count
			p.orig_str = bits[startindex : count+1]
			fmt.Println("Found literal packet ")
			p.toString()
			if isSubPacket {
				pkt.subpackets = append(pkt.subpackets, *p)
			}
		} else {
			fmt.Println("Before :", count, bits)
			// get the length type bit
			p.length_type = bits[6]
			p.length = count
			pkt = getOperatorValues(bits[7:], p)
			fmt.Println("After :", count, p.length, pkt.length, bits)
			count += pkt.length
		}
	}
	return pkt
}

func getOperatorValues(bits string, p *packet) *packet {
	fmt.Println("@@@@@@@@@@@@@@getOperatorValues@@@@@@@@@@@@@@", bits, string(p.length_type))
	// check for length type
	if p.length_type == '0' {
		subpacket_len, _ := strconv.ParseInt(bits[0:15], 2, 64)
		fmt.Println(subpacket_len, bits[15:subpacket_len+15])
		// parse subpackets, could be one or more
		p = parsePacket(bits[15:subpacket_len+15], p, true)
		p.length = int(subpacket_len) + 15
	} else if p.length_type == '1' {
		subpackets_count, _ := strconv.ParseInt(bits[0:11], 2, 64)
		fmt.Println("subpackets_count :", subpackets_count)
		// time to create subpacket under main packet
		subpacket := &packet{}
		p.subpackets = append(p.subpackets, *subpacket)
		// parse subpackets, could be one or more
		p.length += 7
		subpacket = parsePacket(bits[11:], subpacket, true)
		fmt.Println("finished all of the processing...")
		subpacket.toString()
		if int(subpackets_count) == len(p.subpackets) {
			fmt.Printf("Found %d sub-packet(s)\n", len(p.subpackets))
			for _, sub := range p.subpackets {
				p.length += sub.length
			}
		}
		// p.length = int(subpacket_len) + 15
	}
	return p
}

func hexStrToBinaryStr(line string) string {
	var bin bytes.Buffer
	for i := range line {
		bin.WriteString(hexToBinaryDict[(string(line[i]))])
	}
	return bin.String()
}
