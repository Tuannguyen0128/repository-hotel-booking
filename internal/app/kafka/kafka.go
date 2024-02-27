package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"repository-hotel-booking/internal/app/kafka/consumer"
	"repository-hotel-booking/internal/app/kafka/producer"
)

func InitConnection(host string, producerTopic string, consumerTopic string, partition int) (*consumer.Consumer, *producer.Producer) {
	cons, err := kafka.DialLeader(context.Background(), "tcp", host, consumerTopic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	prod, err := kafka.DialLeader(context.Background(), "tcp", host, producerTopic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return &consumer.Consumer{ConsumerConn: cons}, &producer.Producer{
		ProducerConn: prod,
	}
}
