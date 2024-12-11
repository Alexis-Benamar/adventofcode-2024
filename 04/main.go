package main

import (
	_ "embed"
	"fmt"
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

func isOutOfBounds(x, y int) bool {
	return x < 0 || x >= len(lines[0]) || y < 0 || y >= len(lines)
}

func checkPos(charX, charY int, posList [][]CharPos) int {
	var appearances int

	for _, charPosList := range posList {
		isCorrect := true

		for _, charPos := range charPosList {
			posX := charX + charPos.x
			posY := charY + charPos.y

			if (isOutOfBounds(posX, posY) || string(lines[posY][posX]) != charPos.char) {
				isCorrect = false
				break
			}
		}

		if (isCorrect) {
			appearances++
		}
	}

	return appearances
}

func part1() {
	start := time.Now()

	var xmasAppearances int

	for lineId, line := range lines {
		for charId, char := range line {
			if (char != 'X') {
				continue
			}

			xmasAppearances += checkPos(charId, lineId, XMASPositions)
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("part1: %d\n", xmasAppearances)
	fmt.Printf("Execution time %f s\n\n", elapsed.Seconds())
}

func part2() {
	start := time.Now()

	var masAppearances int

	for lineId, line := range lines {
		for charId, char := range line {
			if (char != 'A') {
				continue
			}

			if validAppearances := checkPos(charId, lineId, MASPositions); validAppearances == 2 {
				masAppearances++
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("part2: %d\n", masAppearances)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}

func main() {
	part1()
	part2()
}