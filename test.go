package main

import (
	"fmt"
	rand2 "math/rand"
)

//func main() {
//	// test0001()
//	// rand()
//	// test0001()
//
//	value := 128
//
//	v := value << 8
//	fmt.Println(v)
//
//	//v := value >> 8
//	//fmt.Println(v)
//
//	f := convertTF(129)
//	fmt.Println(f)
//}

func test0001() {
	//port > 0xffff

	// println(0x00000000)

	// imageToTensorTF(10, 10)
}

func rand111() {
	fmt.Println(rand2.Intn(100))
}

//func imageToTensorTF(imageHeight, imageWidth int) {
//	var tfImage [1][][][3]float32
//
//	for j := 0; j < imageHeight; j++ {
//		temp := make([][3]float32, imageWidth)
//		fmt.Println(temp)
//		tfImage[0][0] = append(tfImage[0][0], temp)
//		fmt.Println(tfImage)
//		fmt.Println(tfImage[0])
//		fmt.Println(tfImage[0][0])
//	}
//
//	for i := 0; i < imageWidth; i++ {
//		for j := 0; j < imageHeight; j++ {
//			r, g, b, _ := uint32(1), uint32(2), uint32(3), uint32(4)
//			tfImage[0][j][i][0] = convertTF(r)
//			tfImage[0][j][i][1] = convertTF(g)
//			tfImage[0][j][i][2] = convertTF(b)
//		}
//	}
//
//	fmt.Println(tfImage)
//}

func convertTF(value uint32) float32 {
	return (float32(value>>8) - float32(127.5)) / float32(127.5)
}

//var test_ring *hashring.HashRing
//
//func TestGetHash() {
//
//	arr := []string{"o1"}
//
//	test_ring = hashring.New(arr)
//
//	// add 10 nodes
//	for i := 0; i <= 10; i++ {
//		rk := fmt.Sprintf("online_%d", i)
//		// add RK
//		// RegisterNewNode(rk)
//		// AddNode(rk)
//		test_ring = test_ring.AddNode(rk)
//		// arr = append(arr, rk)
//	}
//
//	fmt.Println(test_ring.Size())
//
//	orders := GetFakeData(100)
//	for _, o := range orders {
//		n, _ := test_ring.GetNode(o.OrderCode)
//		fmt.Println(n)
//	}
//}
