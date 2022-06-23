package messaging

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"github.com/brenno-calado/position_simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/brenno-calado/position_simulator/app/routes"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJSONPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
