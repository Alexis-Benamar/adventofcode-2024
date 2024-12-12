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
	var sum int

	for _, rule := range orderingRules {
		splitRule := strings.Split(rule, "|")

		rulesMap[splitRule[0]] = append(rulesMap[splitRule[0]], splitRule[1])
	}

	for _, update := range pageUpdates {
		var isCorrect bool = true
		pageList := strings.Split(update, ",")

		for i := 1; i < len(pageList); i++ {
			numToCheck := pageList[i]
			numsBefore := pageList[:i]

			for _, num := range numsBefore {
				if slices.Contains(rulesMap[numToCheck], num) {
					isCorrect = false
					break
				}
			}

			if !isCorrect {
				break
			}
		}

		if (isCorrect) {
			midIndexValue, _ := strconv.Atoi(pageList[len(pageList)/2])
			sum += midIndexValue
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", sum)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}