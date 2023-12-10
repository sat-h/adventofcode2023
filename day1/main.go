package main

import (
	"aoc2023/aoclib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var tensTotal, onesTotal int

	for _, line := range aoclib.ConvertToSliceOfString("input.txt") {
		//fmt.Println(line)
		tensReg := regexp.MustCompile(`\d\w*`)
		onesReg := regexp.MustCompile(`\w*\d`)

		tens, _ := strconv.Atoi(string(tensReg.Find([]byte(line))[0]))
		onesLen := len(string(onesReg.Find([]byte(line))))
		ones, _ := strconv.Atoi(string(onesReg.Find([]byte(line))[onesLen-1]))

		tensTotal += tens
		onesTotal += ones
	}

	fmt.Println(tensTotal*10 + onesTotal)
	//55971

	fmt.Println(partTwo())

}

func partTwo() int {
	// 54719
	var tensTotal, onesTotal int

	replacer1 := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	replacer2 := strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")

	for _, line := range aoclib.ConvertToSliceOfString("input_2.txt") {

		reversedLine := aoclib.ReverseString(line)

		reg := regexp.MustCompile(`\d\w*`)

		line = replacer1.Replace(line)
		reversedLine = replacer2.Replace(reversedLine)

		tens, _ := strconv.Atoi(string(reg.Find([]byte(line))[0]))
		ones, _ := strconv.Atoi(string(reg.Find([]byte(reversedLine))[0]))

		tensTotal += tens
		onesTotal += ones
	}

	return (tensTotal*10 + onesTotal)
}
