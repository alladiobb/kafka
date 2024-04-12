package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//chan é um canal que o
	deliveryChannel := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("mensagem", "test", producer, nil, deliveryChannel)

	e := <-deliveryChannel
	msg_return := e.(*kafka.Message)

	if msg_return.TopicPartition.Error != nil {
		fmt.Println("Error ao enviar msg!")
	} else {
		fmt.Println("Mensagem Enviada", msg_return.TopicPartition)
	}

	producer.Flush(1000)

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-kafka-1:9092",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChannel)
	if err != nil {
		return err
	}
	return nil
}
