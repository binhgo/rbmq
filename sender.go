package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func startSender() {

	RbConn := GetRBMQConn()

	fmt.Println("sender connected")

	channel, err := RbConn.Channel()
	failOnError(err, "CHANNEL()")
	// defer channel.Close()

	err = channel.ExchangeDeclare(
		EXCHANGE_NAME,        // name
		EXCHANGE_KIND_DIRECT, // type
		EXCHANGE_DURABLE,     // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	orders := GetFakeData(TEST_MAX_ORDER)
	for _, o := range orders {
		PushToQueue(channel, o)
	}

	fmt.Println("sender stop.")
}

func PushToQueue(channel *amqp.Channel, order FakeData) {

	body, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	consumerName, ok := GetNodeFromRing(order.OrderCode)
	if !ok {
		panic("FUCKKK")
	}

	err = channel.Publish(
		EXCHANGE_NAME,
		consumerName,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "text/plain",
			Body:         body,
		})

	failOnError(err, "Failed to publish a message")

	str := fmt.Sprintf("send to %s order %s", consumerName, order.OrderCode)
	fmt.Println(str)
}
