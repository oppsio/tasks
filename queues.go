package main

import (
	"log"
	"os"

	"encoding/json"
	"github.com/oppsio/tasks/redismq"
)

func print() {
	data, err := json.Marshal(&Settings{})
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	Produce("teams", "Boca Juniors")
	Consume("teams")
}

func Produce(queueName string, payload string) {
	queue := createQueue(queueName)
	err := queue.Put(payload)
	if err != nil {
		log.Fatalln("Cannot connect to Redis")
	}
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
	log.Println("> " + p.Payload)
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
