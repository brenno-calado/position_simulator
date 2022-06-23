package main

import (
	"fmt"
	"log"
	"github.com/brenno-calado/position_simulator/infra/kafka"
	kafka2 "github.com/brenno-calado/position_simulator/app/messaging"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		 go kafka2.Produce(msg)
	}
}