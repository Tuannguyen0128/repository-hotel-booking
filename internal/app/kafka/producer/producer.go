package producer

import (
	"github.com/segmentio/kafka-go"
	"log"
)

func (k *Producer) SendMessage(message []byte) {
	_, err := k.ProducerConn.WriteMessages(
		kafka.Message{Value: message},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

type Producer struct {
	ProducerConn *kafka.Conn
}
