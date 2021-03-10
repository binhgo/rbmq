package main

import (
	"fmt"
	"strings"
)

// test: "aabababaa"
//
// test: ""
func findLongestPalindromic(input string) {

	m := make(map[int]int)

	for i, _ := range input {

		if i == 0 || i == len(input)-1 {
			continue
		}

		max := 0
		if i < len(input)/2 {
			max = i
		} else {
			max = len(input) - 1 - i
		}

		for j := 1; j <= max; j++ {
			if input[i-j] == input[i+j] {
				m[i]++
			}
		}

	}

	max := 0
	index := 0
	for i, v := range m {
		if v > max {
			max = v
			index = i
		}
	}

	fmt.Println(index)
	fmt.Println(m[index])
	str := input[index+1 : index+1+m[index]]
	fmt.Println(strings.ReplaceAll(str, "#", ""))
}

func main() {
	// input := "abaaba"
	input := "sqabawabaqs"

	newInput := string(input[0])

	for i, v := range input {

		if i == 0 {
			continue
		}

		newInput = fmt.Sprintf("%s#%s", newInput, string(v))
	}

	fmt.Println(newInput)

	findLongestPalindromic(newInput)
}
