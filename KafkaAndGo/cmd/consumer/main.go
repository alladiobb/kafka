package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafkaandgo-kafka-1:9092",
		"client.id":         "goapp2-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("error Consumer", err.Error())
	}
	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)
	for {
		mensagem, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(mensagem.Value), mensagem.TopicPartition)
		}
	}
}
