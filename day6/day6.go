package main

import (
	//"aoc2023/aoclib"
	"aoc2023/aoclib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partTwo("input.txt"))
}

func partOne(filename string) int {

	var totalWays int

	lines := aoclib.ConvertToSliceOfSliceOfString(filename)
	time := lines[0]
	distance := lines[1]

	for i := 1; i < len(time); i++ {
		t, _ := strconv.ParseFloat(time[i], 32)
		d, _ := strconv.ParseFloat(distance[i], 32)

		bigger, smaller := solveQuadraticEquation(t, d)
		ways := int(bigger - smaller + 1)
		if i == 1 {
			totalWays = ways
		} else {
			totalWays *= ways
		}

	}
	return totalWays
}

func partTwo(filename string) int {
	lines := aoclib.ConvertToSliceOfSliceOfString(filename)

	time := strings.Join(lines[0][1:], "")
	distance := strings.Join(lines[1][1:], "")
	t, _ := strconv.ParseFloat(time, 64)
	d, _ := strconv.ParseFloat(distance, 64)
	fmt.Println(t, d)

	bigger, smaller := solveQuadraticEquation(t, d)

	return int(bigger - smaller + 1) // 24655065 too low
}

func solveQuadraticEquation(t float64, d float64) (float64, float64) {

	x1 := (t + math.Sqrt(math.Pow(t, 2)-4*d)) / 2
	x2 := (t - math.Sqrt(math.Pow(t, 2)-4*d)) / 2

	fmt.Println(x1, x2)

	smaller := x1
	bigger := x2

	if x1 > x2 {
		smaller = x2
		bigger = x1
	}
	return math.Floor(bigger), math.Ceil(smaller)
}
