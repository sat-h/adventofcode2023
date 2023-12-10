package main

import (
	"aoc2023/aoclib"
	"fmt"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {

	var total int

	s := aoclib.ConvertToSliceOfSliceOfString("input.txt")

	for _, v := range s {
		s := aoclib.ConvertToSliceOfInt(v)
		total += findNextValue(s) + s[len(s)-1]
	}

	fmt.Println(total)
}

func partTwo() {
	
	var total int

	s := aoclib.ConvertToSliceOfSliceOfString("input.txt")

	for _, v := range s { // s: original sequence
		s := aoclib.ConvertToSliceOfInt(v)
		total += s[0] - findPreviousValue(s)
	}
	fmt.Println(total)
}

func findNextValue(s []int) int { // for each line of input

	var nextValue int
	var sequences [][]int

	num := len(s) - 1
	for i := 0; i < num; i++ {
		newSequence := generateNewSequence(s)
		sequences = append(sequences, newSequence)
		s = newSequence
	}

	for _, s := range sequences {
		nextValue += s[len(s)-1]
	}

	return nextValue
}

func findPreviousValue(s []int) int {

	var sequences [][]int
	var seed int

	num := len(s) - 1
	for i := 0; i < num; i++ {
		newSequence := generateNewSequence(s)
		sequences = append(sequences, newSequence)
		s = newSequence
	}
	
	for i := len(sequences) - 1; i >= 0; i-- {
		if sequences[i][0] != 0 {
			sequences[i] = append([]int{sequences[i][0]}, sequences[i]...)
			seed = i
			break
		}
	}

	for i := seed; i > 0; i-- {
		sequences[i-1] = append([]int{sequences[i-1][0]-sequences[i][0]}, sequences[i-1]...)
	}

	return sequences[0][0]
}

func generateNewSequence(s []int) []int {
	var nextLine []int
	for i := 0; i < len(s)-1; i++ {
		nextLine = append(nextLine, s[i+1]-s[i])
	}
	return nextLine
}
