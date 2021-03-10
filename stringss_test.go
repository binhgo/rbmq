package main

import (
	"fmt"
	"testing"
)

func BenchmarkAddNodeToRing(b *testing.B) {

	for i := 0; i < b.N; i++ {
		AddNodeToRing("AAADSDE#$#FDFDF")
	}

}

func BenchmarkPlusAB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PlusAB("AA", "BBBB")
	}
}

func TestPlusAB(t *testing.T) {

	result := PlusAB("AA", "B")
	if result != "AA0B" {
		fmt.Println(result)
		t.Fail()
	}
}

func TestPlusAB2(t *testing.T) {
	r2 := PlusAB("AA", "BB")
	if r2 != "AABB" {
		t.Fail()
	}
}
