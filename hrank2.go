package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func findMax(b []int32) int32 {

	max := b[0]
	for _, v := range b {
		if v > max {
			max = v
		}
	}

	return max
}

func isFactorOfAll(a []int32, num int32) bool {

	for _, v := range a {
		if num%v != 0 {
			return false
		}
	}

	return true
}

func isOkWithB(b []int32, num int32) bool {
	for _, v := range b {
		if v%num != 0 {
			return false
		}
	}

	return true
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here

	// find max number in B[] called maxB
	maxB := findMax(b)

	// find the max number = maxB/2
	maxC := maxB / 2

	total := int32(0)
	// loop
	for i := 1; i <= int(maxC); i++ {

		ok := isOkWithB(b, int32(i))
		if ok {
			ok := isFactorOfAll(a, int32(i))
			if ok {
				total++
			}
		}
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
