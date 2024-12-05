package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
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

func areLevelsSafe(levels []string) bool {
	for i := 0; i < len(levels)-1; i++ {
		var num1, num2 int

		num1, _ = strconv.Atoi(levels[i])
		num2, _ = strconv.Atoi(levels[i+1])

		if i > 0 {
			num0, _ := strconv.Atoi(levels[i-1])

			if !((num0 < num1 && num1 < num2) || (num0 > num1 && num1 > num2)) {
				return false
			}
		}

		if num1 == num2 {
			return false
		}

		if math.Abs(float64(num2 - num1)) > 3 {
			return false
		}

	}

	return true
}

func main() {
	start := time.Now()

	var safeLevelsNb int
	reNum := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		levels := reNum.FindAllString(line, -1)

		if areLevelsSafe(levels) {
			safeLevelsNb++
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", safeLevelsNb)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}