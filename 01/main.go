package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var data string
var lines []string

func init() {
	data = strings.TrimRight(data, "\n")
	if len(data) == 0 {
		panic("empty data file")
	}

	lines = strings.Split(data, "\n")
}

func main() {
	start := time.Now()

	var sumDistances int
	var similarityScore int
	var leftNum string
	var rightNum string
	leftList := make([]string, len(lines))
	rightList := make([]string, len(lines))

	for i, line := range lines {
		_, _ = fmt.Sscanf(line, "%s   %s", &leftNum, &rightNum)

		leftList[i] = leftNum
		rightList[i] = rightNum
	}

	sort.Strings(leftList)
	sort.Strings(rightList)
	rightListAsString := strings.Join(rightList, " ")

	for i := range lines {
		leftAsNumber, _ := strconv.Atoi(leftList[i])
		rightAsNumber, _ := strconv.Atoi(rightList[i])

		if (rightAsNumber > leftAsNumber) {
			sumDistances += (rightAsNumber - leftAsNumber)
		} else {
			sumDistances += (leftAsNumber - rightAsNumber)
		}

		regex := regexp.MustCompile(leftList[i])
		match := regex.FindAllStringIndex(rightListAsString, -1)

		similarityScore += leftAsNumber * len(match)
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", sumDistances)
	fmt.Printf("part2: %d\n", similarityScore)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}