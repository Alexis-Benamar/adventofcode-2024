package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var rawData string
var data string

func init() {
	rawData = strings.TrimRight(rawData, "\n")
	if len(rawData) == 0 {
		panic("empty data file")
	}

	data = strings.Split(rawData, "\n\n")[0]
}

func checkMemory(data string, part int8) int {
	var sum int
	reAll := regexp.MustCompile(`(don't\(\))|(do\(\)|(mul\((\d{1,3}),(\d{1,3})\)))`)

	var ignoreMul bool
	ignoreMulMap := map[string]bool{"do()": false,"don't()": true}

	allMatches := reAll.FindAllStringSubmatch(data, -1)

	for _, match := range allMatches {
		if (part == 2) {
			if !strings.HasPrefix(match[0], "mul") {
				ignoreMul = ignoreMulMap[match[0]]
				continue
			}

			if ignoreMul {
				continue
			}
		}

		num1, _ := strconv.Atoi(match[4])
		num2, _ := strconv.Atoi(match[5])
		sum += num1 * num2
	}

	return sum
}

func main() {
	start := time.Now()

	part1 := checkMemory(data, 1)
	part2 := checkMemory(data, 2)

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}