package main

import (
	"fmt"

	"github.com/serialx/hashring"
)

var ring *hashring.HashRing

func GetNodeFromRing(key string) (string, bool) {

	if ring == nil {
		ring = hashring.New(GetAllConsumerId())
	}

	return ring.GetNode(key)
}

func AddNodeToRing(node string) {

	if ring == nil {
		ring = hashring.New(GetAllConsumerId())
	}

	fmt.Println("added node: " + node)
	ring = ring.AddNode(node)
}
