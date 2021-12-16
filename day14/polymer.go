package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Ramkrao/advent/utils"
)

var startTime = time.Now().UnixMilli()

func main() {
	lines := utils.ReadArrayFromFile("day14/input.txt")

	template := make(map[string]int64)
	rules := make(map[string]string)
	for i := range lines {
		if !strings.Contains(lines[i], "->") && len(lines[i]) > 0 {
			for j := 0; j < len(lines[i])-1; j++ {
				template[string(lines[i][j])+string(lines[i][j+1])] += 1
			}
		} else if strings.Contains(lines[i], "->") {
			rule := strings.Split(lines[i], " -> ")
			rules[rule[0]] = rule[1]
		}
	}
	fmt.Println(template, rules)

	for i := 0; i < 40; i++ {
		copyTemplate := make(map[string]int64)
		for k, v := range template {
			element, prs := rules[string(k)]
			if prs {
				copyTemplate[string(k[0])+element] += v
				copyTemplate[element+string(k[1])] += v
			}
		}
		template = copyTemplate
		fmt.Println("After step ", i, ": ", template)
	}
	results := make(map[string]int64)
	for k, v := range template {
		fmt.Println(k, v)
		results[string(k[0])] += v
	}
	fmt.Println("calculating totals ", results)
	var total []int64
	for _, v := range results {
		total = append(total, v)
	}
	sort.Slice(total, func(i, j int) bool { return total[i] < total[j] })
	sum := total[len(total)-1] - total[0]
	fmt.Println(sum, total, total[0], total[len(total)-1])
}
