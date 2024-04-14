package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	//chan é um canal que o Kafka usa para pegar o retorno
	deliveryChannel := make(chan kafka.Event)

	producer := NewKafkaProducer()

	// Essa linha da "KEY" Não deixa mandar para outras partições. No caso não vai ter problema da ordem de envio
	// key := []byte("1")

	Publish("mensagem", "teste", producer, nil, deliveryChannel)

	//Forma Sincrona:
	// e := <-deliveryChannel
	// msg_return := e.(*kafka.Message)

	// if msg_return.TopicPartition.Error != nil {
	// 	fmt.Println("Error ao enviar msg!")
	// } else {
	// 	fmt.Println("Mensagem Enviada Sincrona: ", msg_return.TopicPartition)
	// }

	// Joga em uma thread de forma assincrona - Trava o terminal
	// DeliveryReport(deliveryChannel)

	// Joga em uma thread de forma assincrona em outra thread - Não trava o terminal
	go DeliveryReport(deliveryChannel)

	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafkaandgo-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		//"all" -> aguardar todos os sync brokers receberem as msg
		//"0" -> sem retorno
		//"1" ->
		"enable.idempotence": "true",
		//"true" -> Certeza que a mensagem for entregue apenas uma vez
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

	// Sem usar o DeliveryChannel
	// err := producer.Produce(message, nil)
	// if err != nil {
	// 	return err
	// }
	// return nil

	//Usando o DeliveryChannel
	err := producer.Produce(message, deliveryChannel)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChannel chan kafka.Event) {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Error ao enviar msg Assincrona!")
			} else {
				fmt.Println("Mensagem Assincrona Enviada: ", ev.TopicPartition)
				//Criar logs no sistema
			}
		}
	}
}
