package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if x1 == x2 {
		if v1 != v2 {
			return "NO"
		} else {
			return "YES"
		}
	}

	//0  3 4  2 :4 diff:4 step:1
	//3  3 6  2 :3 diff:3 step:1
	//6  3 8  2 :2 diff:2 step:1
	//9  3 10 2 :1
	//12 3 12 2 :0

	//0  4 4  2 :4 diff:4 step:2
	//4  4 6  2 :2 diff:2 step:2
	//8  4 8  2 :0

	//0  5 4  2 :4 diff:4 step:3
	//5  5 6  2 :1 diff:1 step:3
	//10 5 8  2 :-2
	//15 5 10 2 :-5

	if x1 > x2 {

		if v1 == v2 {
			return "NO"
		}

		if v2 < v1 {
			return "NO"
		}

		diff := x1 - x2
		step := v2 - v1

		if diff%step == 0 {
			return "YES"
		}

		return "NO"
	}

	if x2 > x1 {

		if v1 == v2 {
			return "NO"
		}

		if v1 < v2 {
			return "NO"
		}

		diff := x2 - x1
		step := v1 - v2

		if diff%step == 0 {
			return "YES"
		}

		return "NO"

	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	x1V1X2V2 := strings.Split(readLine(reader), " ")

	x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
	checkError(err)
	x1 := int32(x1Temp)

	v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
	checkError(err)
	v1 := int32(v1Temp)

	x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
	checkError(err)
	x2 := int32(x2Temp)

	v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
	checkError(err)
	v2 := int32(v2Temp)

	result := kangaroo(x1, v1, x2, v2)

	fmt.Fprintf(writer, "%s\n", result)

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
