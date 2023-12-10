package aoclib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ConvertToSliceOfString(fileName string) []string {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\r\n")
}

func ConvertToSliceOfSliceOfString(fileName string) [][]string {

	var sliceOfSliceofString [][]string

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	sliceOfString := strings.Split(string(data), "\r\n")

	for _, s := range sliceOfString {
		splittedByWhitespaceSlice := strings.Split(s, " ")
		whiteSpaceRemovedSlice := RemoveElementFromSlice(splittedByWhitespaceSlice, "")
		sliceOfSliceofString = append(sliceOfSliceofString, whiteSpaceRemovedSlice)
	}

	return sliceOfSliceofString
}

func ConvertToSliceOfInt(s []string) []int {
	var sliceOfInt []int
	for _, v := range s {
		v, _ := strconv.Atoi(v)
		sliceOfInt = append(sliceOfInt, v)
	}
	return sliceOfInt
}

func RemoveElementFromSlice[T comparable](s []T, item T) []T {
	out := make([]T, 0)
	for _, element := range s {
		if element != item {
			out = append(out, element)
		}
	}
	return out
}

func ReverseString(s string) string {
	var newString string
	for i := len(s) - 1; i >= 0; i-- {
		newString += string(s[i])
	}
	return newString
}

func FindMin(s []int) int {
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func FindMax(s []int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
