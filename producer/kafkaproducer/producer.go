package kafkaproducer

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/ornellast/bucketeer/producer/models"
	"github.com/segmentio/kafka-go"
)

var (
	topicName string = "first-bucketeer"
	brokers          = []string{"host.docker.internal:9092"}
	// brokers          = []string{"kafka1:9092"}
)

func SendToKafka(item *models.Item) {
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

	itemByteArray, err := json.Marshal(item)

	if err != nil {
		log.Printf("Error when marshaling Item: %s\n\t%v", err, item)
		return
	}

	err = producer.WriteMessages(context.Background(),
		kafka.Message{

			Key:   []byte(strconv.Itoa(item.ID)),
			Value: itemByteArray,
		})

	if err != nil {
		log.Println("Failed to write messages: ", err)
	}
}
