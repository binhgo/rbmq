package main

import (
	"fmt"

	"github.com/serialx/hashring"
)

var test_ring *hashring.HashRing

func TestGetHash() {

	arr := []string{"o1"}

	test_ring = hashring.New(arr)

	// add 10 nodes
	for i := 0; i <= 10; i++ {
		rk := fmt.Sprintf("online_%d", i)
		// add RK
		// RegisterNewNode(rk)
		// AddNode(rk)
		test_ring = test_ring.AddNode(rk)
		// arr = append(arr, rk)
	}

	fmt.Println(test_ring.Size())

	orders := GetFakeData(100)
	for _, o := range orders {
		n, _ := test_ring.GetNode(o.OrderCode)
		fmt.Println(n)
	}
}
