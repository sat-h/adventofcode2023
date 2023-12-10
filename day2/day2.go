package main

import (
	aoclib "aoc2023/aoclib"
	"fmt"

	"strconv"
	"strings"

	//"github.com/yuin/goldmark/util"
)

func main() {

	fmt.Println(partOne("input_1.txt"))
	fmt.Println(partTwo("input_1.txt"))

}

func partOne(filename string) int {
	// 2113 too high
	for _, line := range aoclib.ConvertToSliceOfString(filename) {
		fmt.Printf("%#v\n", line)
	}

	inputLines := aoclib.ConvertToSliceOfString(filename)
	fmt.Printf("%#v\n", inputLines)
	var ct int
	for i := 1; i <= len(inputLines); i++ {
		fmt.Println(i)
		ct += i
	}

	fmt.Println(ct)
	for _, line := range inputLines {
		line = strings.Replace(line, "\r", "", -1)
		splittedLine := strings.Split(line, " ")
		rawId := strings.Replace(splittedLine[1], ":", "", 1)
		id, _ := strconv.Atoi(rawId)

	L:
		for i := 3; i <= len(splittedLine); i += 2 {

			rawClour := strings.Replace(splittedLine[i], ",", "", -1)
			rawClour = strings.Replace(rawClour, ";", "", -1)
			colour := rawClour
			num, _ := strconv.Atoi(string(splittedLine[i-1]))

			switch colour {
			case "green":
				if num > 13 {
					ct -= id
					break L
				}
			case "red":
				if num > 12 {
					ct -= id
					break L
				}
			case "blue":
				if num > 14 {
					ct -= id
					break L
				}
			}
		}
	}
	return ct
}

func partTwo(filename string) int {
	var sum int
	inputLines := aoclib.ConvertToSliceOfString(filename)
	replacer := strings.NewReplacer(":", "", ",", "", ";", "", "\r", "")
	for _, line := range inputLines {
		line = replacer.Replace(line)
		splittedLine := strings.Split(line, " ")
		id := splittedLine[1]
		fmt.Println(id, line)
		var green, red, blue []int
		for i := 3; i <= len(splittedLine); i += 2 {

			colour := splittedLine[i]
			num, _ := strconv.Atoi(splittedLine[i-1])
			fmt.Printf("colour: %#v, num: %v\n", colour, num)

			switch colour {
			case "green":
				green = append(green, num)
			case "red":
				red = append(red, num)
			case "blue":
				blue = append(blue, num)
			}
		}

	
		r := aoclib.FindMax(red)
		b := aoclib.FindMax(blue)
		g := aoclib.FindMax(green)

		power := r * b * g
		sum += power

	}
	return sum
}
