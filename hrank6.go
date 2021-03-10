package main

import "fmt"

/*
 * clockwise rotate 顺时针旋转
 * first reverse up to down,
	then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
*/

func rotateMatrix(input [][]int) {

	// first reverse up to down
	r := reverseUpDown(input)

	// then swap the symmetry
	l := len(input)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			ij, ji := r[i][j], r[j][i]
			r[i][j] = ji
			r[j][i] = ij
		}
	}

	fmt.Println(r)
}

func reverseUpDown(input [][]int) [][]int {

	l := len(input)
	result := make([][]int, l)

	for i, _ := range result {
		result[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		result[l-i-1] = input[i]
	}
	
	return result
}

func main() {

	input := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	fmt.Println(input)

	rotateMatrix(input)
}
