package main

//func main() {
//	rb := ringbuffer.New(1024)
//
//	// write
//	rb.Write([]byte("abcd"))
//	fmt.Println(rb.Length())
//	fmt.Println(rb.Free())
//
//	// read
//	buf := make([]byte, 8)
//	rb.Read(buf)
//	fmt.Println(string(buf))
//
//	fmt.Println(Factorial(50))
//}

func Factorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}
