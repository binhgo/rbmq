package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

var rbmqConn *amqp.Connection

func GetRBMQConn() *amqp.Connection {

	// if != and status != close
	if rbmqConn != nil && rbmqConn.IsClosed() == false {
		return rbmqConn
	}

	conn, err := amqp.Dial(RBMQ_CONN_STRING)
	failOnError(err, "Failed to connect to RabbitMQ")

	rbmqConn = conn

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
