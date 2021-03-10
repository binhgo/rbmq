package main

import "fmt"

// abcd

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

// 3:4 = d
// MCMXCIV

// V : 5
// I : 1

func main() {
	input := "MCMXCIV" // 1994
	rs := romanToInt(input)
	fmt.Println(rs)
}

func romanToInt(input string) int {

	total := 0

	mm := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	l := len(input)
	lastInt := 0
	for i := l - 1; i >= 0; i-- {
		c := input[i]

		digit := mm[string(c)]
		if digit > lastInt {
			total = total + digit
		} else {
			total = total - digit
		}

		lastInt = digit
	}

	return total
}
