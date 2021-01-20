package main

import (
	"time"
)

var consumerList []Consumer

type Consumer struct {
	ConsumerId string
	LastPing   time.Time
}

func GetAllConsumerId() []string {
	var arr []string
	for _, v := range consumerList {
		if v.LastPing.After(time.Now().Add(-10 * time.Second)) {
			arr = append(arr, v.ConsumerId)
		}
	}

	return arr
}

func GetAllConsumers() []Consumer {
	return consumerList
}

func RegisterNewConsumer(consumerName string) error {

	csm := Consumer{
		ConsumerId: consumerName,
		LastPing:   time.Now(),
	}

	consumerList = append(consumerList, csm)

	// add to hash ring
	AddNodeToRing(consumerName)

	return nil
}

func RemoveConsumer(consumerName string) error {

	for i, v := range consumerList {
		if v.ConsumerId == consumerName {
			consumerList[i].ConsumerId = ""
		}
	}

	return nil
}

func ConsumerPingRegistry(consumerName string) error {

	for i, v := range consumerList {
		if v.ConsumerId == consumerName {
			consumerList[i].LastPing = time.Now()
		}
	}

	return nil
}
