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

// TODO: move to helper file
func arrayCopy(orig []string) []string {
	newArray := make([]string, len(orig))
	copy(newArray, orig[:])
	return newArray
}

func checkLevels(levels []string) (bool, int) {
	for i := 0; i < len(levels)-1; i++ {
		var num1, num2 int

		num1, _ = strconv.Atoi(levels[i])
		num2, _ = strconv.Atoi(levels[i+1])

		if i > 0 {
			num0, _ := strconv.Atoi(levels[i-1])

			if !((num0 < num1 && num1 < num2) || (num0 > num1 && num1 > num2)) {
				return false, i
			}
		}

		if num1 == num2 {
			return false, i
		}

		if math.Abs(float64(num2 - num1)) > 3 {
			return false, i
		}
	}

	return true, -1
}

func main() {
	start := time.Now()

	var safeLevelsNb, safeishLevelsNb int
	reNum := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		levels := reNum.FindAllString(line, -1)

		levelIsSafe, i := checkLevels(levels)

		if (levelIsSafe) {
			safeLevelsNb++
		} else  {
			for j := -1; j <= 1; j++ {
				if i+j < 0 {
					continue
				}

				levelsCopy := arrayCopy(levels)
				levelsCopy = append(levelsCopy[:i+j], levelsCopy[i+j+1:]...)

				isSafeWithJRemoved, _ := checkLevels(levelsCopy)

				if isSafeWithJRemoved {
					safeishLevelsNb++
					break
				}
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", safeLevelsNb)
	fmt.Printf("part2: %d\n", safeLevelsNb + safeishLevelsNb)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}