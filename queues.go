package main

import (
	"log"
	"os"

	"github.com/adjust/redismq"
)

func main() {
	// Produce("teams", "Boca Juniors")
	// Consume("teams")
}

func Produce(queueName string, payload string) {
	queue := createQueue(queueName)
	queue.Put(payload)
	log.Println("produced:", payload)
}

func Consume(queueName string) {
	queue := createQueue(queueName)
	consumer, err := queue.AddConsumer(queueName)
	if err != nil {
		os.Exit(0)
	}

	p, err := consumer.NoWaitGet()
	if err != nil {
		log.Println(err)
	}
	if p == nil {
		os.Exit(0) // Nothing to consume
	}
	log.Println(p.CreatedAt)
	log.Println(p.Payload)
	err = p.Ack()
	if err != nil {
		log.Println(err)
	}
}

func createQueue(queueName string) *redismq.Queue {
	redisAddr := os.Getenv("REDIS_PORT_6379_TCP_ADDR")
	redisPort := os.Getenv("REDIS_PORT_6379_TCP_PORT")
	if redisAddr == "" {
		redisAddr = "localhost"
	}
	if redisPort == "" {
		redisPort = "6379"
	}
	return redismq.CreateQueue(redisAddr, redisPort, "", 9, queueName)
}
