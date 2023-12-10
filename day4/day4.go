package main

import (
	aoclib "aoc2023/aoclib"
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

var cardMap = map[int]*card{}

type card struct {
	matching int
	copies   int
}

func main() {

	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {

	var ct int

	for j, line := range aoclib.ConvertToSliceOfString("input.txt") {

		line = strings.Replace(line, ":", "", -1) 
		newLine := strings.Split(line, " | ")    

		var offSet int
		if j < 9 {
			offSet = 4
		} else if j < 98 {
			offSet = 3
		} else {
			offSet = 2
		}
		winNums := strings.Split(newLine[0], " ")[offSet:]
		yourNums := strings.Split(newLine[1], " ")

		var subCt int
		for i := 0; i < len(winNums); i++ {
			winNum := winNums[i]
			if winNum != "" {
				if slices.Contains(yourNums, winNum) {
					subCt++
				}
			}
		}
		cardMap[j+1] = &card{matching: subCt, copies: 1}
		ct += int(math.Pow(2, float64(subCt)-1))
		subCt = 0
	}

	return ct // 20750 too high
}

func partTwo() int {

	for i := 1; i <= len(cardMap); i++ {
		c := cardMap[i]
		for k := 0; k < c.copies; k++ {
			for j := 1; j <= c.matching; j++ {
				cardNum := i + j
				c2 := cardMap[cardNum]
				c2.copies += 1
			}
		}
	}

	p := 0
	for _, card := range cardMap {
		p += card.copies
	}

	return p

	// 5761573 too high
}
