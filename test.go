package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	rs := TempSock(1234)
	println(rs)
	println(os.TempDir())
}

func TempSock(totalDuration float64) string {
	// serve

	rand.Seed(time.Now().Unix())
	sockFileName := path.Join(os.TempDir(), fmt.Sprintf("%d_sock", rand.Int()))
	l, err := net.Listen("unix", sockFileName)
	if err != nil {
		panic(err)
	}

	go func() {
		re := regexp.MustCompile(`out_time_ms=(\d+)`)
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		buf := make([]byte, 16)
		data := ""
		progress := ""
		for {
			_, err := fd.Read(buf)
			if err != nil {
				return
			}
			data += string(buf)
			a := re.FindAllStringSubmatch(data, -1)
			cp := ""
			if len(a) > 0 && len(a[len(a)-1]) > 0 {
				c, _ := strconv.Atoi(a[len(a)-1][len(a[len(a)-1])-1])
				cp = fmt.Sprintf("%.2f", float64(c)/totalDuration/1000000)
			}
			if strings.Contains(data, "progress=end") {
				cp = "done"
			}
			if cp == "" {
				cp = ".0"
			}
			if cp != progress {
				progress = cp
				fmt.Println("progress: ", progress)
			}
		}
	}()

	return sockFileName
}

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
	// fmt.Println(crand.Intn(100))
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
