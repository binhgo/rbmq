package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "abcabc"
	fmt.Println(isBalanceStr(s1))
}

func PlusAB(a string, b string) string {

	time.Sleep(100 * time.Nanosecond)

	l := 0
	needtofill := 0

	A := ""
	B := ""

	if len(a) > len(b) {
		l = len(a)
		needtofill = len(a) - len(b)

		newB := make([]byte, l)
		for i := 0; i < l; i++ {
			if i < needtofill {
				newB[i] = '0'
				continue
			}

			newB[i] = b[i-needtofill]
		}

		A = a
		B = string(newB)

	} else {
		l = len(b)
		needtofill = len(b) - len(a)

		newA := make([]byte, l)
		for i := 0; i < l; i++ {
			if i < needtofill {
				newA[i] = 0
				continue
			}

			newA[i] = a[i-needtofill]
		}

		A = string(newA)
		B = b
	}

	return A + B
}

func isBalanceStr(data string) bool {

	l := len(data)
	if l%2 != 0 {
		return false
	}

	var left []int32
	var right []int32

	// 10 / 2 = 5
	// 0 1 2 3 4 5 6 7 8 9

	for i, v := range data {

		if i >= l/2 {
			// right
			right = append(right, v)

		} else {
			// left
			left = append(left, v)
		}

	}

	for i, v := range left {

		idx := l/2 - 1 - i
		if v != right[idx] {
			return false
		}
	}

	return true
}
