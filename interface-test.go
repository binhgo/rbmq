package main

import (
	"fmt"
)

type IReader interface {
	Read(p []byte) error
}

//func ReadData(reader IReader) error {
//	return reader.Read()
//}

type MyType struct {
	Name string
	Data string
}

func (mt *MyType) Write(p []byte) (n int, err error) {

	for i, v := range p {
		fmt.Printf("i: %d v: %d \n", i, v)
	}

	mt.Data = string(p)
	return len(p), nil
}

func (mt *MyType) Read(b []byte) (n int, err error) {
	l := len(b)
	copy(b, mt.Data[0:l])
	return l, nil
}

//func main() {
//	// var buf bytes.Buffer
//	//fmt.Fprintf(&buf, "Size: %d MB.", 85)
//	//s := buf.String() // s == "Size: 85 MB."
//	//fmt.Fprintf()
//
//	var myType MyType
//	fmt.Fprintf(&myType, "Size: %d MB.S", 85)
//	fmt.Println(myType.Data)
//
//	//r := strings.NewReader("abcde")
//	//buf := make([]byte, 4)
//	//for {
//	//	n, err := r.Read(buf)
//	//	fmt.Println(n, err, buf[:n])
//	//	if err == io.EOF {
//	//		break
//	//	}
//	//}
//
//	fmt.Println("READDDD")
//	buf := make([]byte, 4)
//	n, err := myType.Read(buf)
//	fmt.Println(n, err, buf[:n])
//}
