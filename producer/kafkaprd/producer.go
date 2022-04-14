package kafkaprd

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ornellast/bookstore/producer/commons"
	"github.com/segmentio/kafka-go"
)

var (
	// topicName string = "first-bookstore"
	brokers = []string{"host.docker.internal:9092"}
	// brokers          = []string{"kafka1:9092"}
)

func NewEvent(topicName string, msg commons.Identifier) error {
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topicName,
		// Balancer: &kafka.Murmur2Balancer{},
	})

	log.Printf("Producer Created\n\t%v\n\n\n", producer)
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close the Writer: ", err)
		}
	}()

	msgByteArray, err := json.Marshal(msg)

	if err != nil {
		log.Printf("Error when marshaling Book: %s\n\t%v", err, msg)
		return err
	}

	err = producer.WriteMessages(context.Background(),
		kafka.Message{

			Key:   []byte(msg.Id()),
			Value: msgByteArray,
		})

	if err != nil {
		log.Println("Failed to write messages: ", err)
		return err
	}

	return nil
}
