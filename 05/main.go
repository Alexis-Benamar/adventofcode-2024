package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var instructions []string

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	instructions = strings.Split(data, "\n\n")
}

func main() {
	start := time.Now()

	rulesMap := make(map[string][]string)
	orderingRules := strings.Split(instructions[0], "\n")
	pageUpdates := strings.Split(instructions[1], "\n")
	incorrectPagesUpdates := make([][]string, 0)
	var sum, fixedUpdatesSum int

	for _, rule := range orderingRules {
		splitRule := strings.Split(rule, "|")

		rulesMap[splitRule[0]] = append(rulesMap[splitRule[0]], splitRule[1])
	}

	// part 1
	for _, update := range pageUpdates {
		var isCorrect bool = true
		pageList := strings.Split(update, ",")

		for i := 1; i < len(pageList); i++ {
			numToCheck := pageList[i]
			numBefore := pageList[i-1]

			if slices.Contains(rulesMap[numToCheck], numBefore) {
				isCorrect = false
				incorrectPagesUpdates = append(incorrectPagesUpdates, pageList)
				break
			}
		}


		if (isCorrect) {
			midIndexValue, _ := strconv.Atoi(pageList[len(pageList)/2])
			sum += midIndexValue
		}
	}

	// part 2
	// sort function taken from here https://github.com/mnml/aoc/blob/main/2024/05/1.go#L15
	compare := func(a, b string) int {
		for _, rule := range orderingRules {
			if rule := strings.Split(rule, "|"); rule[0] == a && rule[1] == b {
				return -1
			}
		}
		return 0
	}

	for _, wrongUpdate := range incorrectPagesUpdates {
		slices.SortFunc(wrongUpdate, compare)
		midIndexValue, _ := strconv.Atoi(wrongUpdate[len(wrongUpdate)/2])
		fixedUpdatesSum += midIndexValue
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", sum)
	fmt.Printf("part2: %d\n", fixedUpdatesSum)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}