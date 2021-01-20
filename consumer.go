package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func startConsumer(consumerCount int) {

	RbConn := GetRBMQConn()

	fmt.Println("consumer connected")
	// defer RbConn.Close()

	channel, err := RbConn.Channel()
	failOnError(err, "RbConn.Channel()")
	// defer channel.Close()

	err = channel.ExchangeDeclare(
		EXCHANGE_NAME,
		EXCHANGE_KIND_DIRECT,
		EXCHANGE_DURABLE,
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	for i := 0; i <= consumerCount; i++ {

		consumerName := fmt.Sprintf("online_%d", i)

		q, err := channel.QueueDeclare(
			consumerName,     // name
			EXCHANGE_DURABLE, // durable
			false,            // delete when unused
			true,             // exclusive
			false,            // no-wait
			nil,              // arguments
		)

		err = channel.QueueBind(
			q.Name,        // queue name
			q.Name,        // routing key
			EXCHANGE_NAME, // exchange
			false,
			nil)

		failOnError(err, "Failed to bind a queue")

		// add RK
		err = RegisterNewConsumer(consumerName)
		failOnError(err, "Failed to register a consumer")

		msgs, _ := channel.Consume(
			q.Name,
			consumerName,
			true,
			false,
			false,
			false,
			nil,
		)

		go ConsumeData(msgs, consumerName)
	}
}

func ConsumeData(msgs <-chan amqp.Delivery, consumerName string) {

	go startPingWorker(consumerName, 5) // 5s

	fmt.Println(fmt.Sprintf("Consumer: %s waiting", consumerName))

	for d := range msgs {
		str := fmt.Sprintf("consumer: %s, data: %s", d.RoutingKey, d.Body)
		fmt.Println(str)
	}
}

func startPingWorker(consumerName string, interval time.Duration) {
	ticker := time.NewTicker(interval * time.Second)
	for range ticker.C {
		ConsumerPingRegistry(consumerName)
	}
}
