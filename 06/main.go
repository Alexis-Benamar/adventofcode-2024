package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed example.txt
var area string

func init() {
	area = strings.TrimRight(area, "\n")
	if len(area) == 0 {
		panic("empty data file")
	}
}

func main() {
	start := time.Now()

	totalLength := len(area)
	lineLength := len(strings.Split(area, "\n")[0])

	fmt.Println(totalLength, lineLength)

	elapsed := time.Since(start)
	fmt.Printf("Execution time %f s\n", elapsed.Seconds())
}