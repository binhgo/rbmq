package main

import "fmt"

type FakeData struct {
	OrderCode string
	Data      string
}

func GetFakeData(l int) []FakeData {

	var arr []FakeData

	for i := 1; i <= l; i++ {
		oc := fmt.Sprintf("order_%d", i)

		f1 := FakeData{
			OrderCode: oc,
			Data:      oc,
		}

		arr = append(arr, f1)
	}

	return arr
}
