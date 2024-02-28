package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"repository-hotel-booking/internal/app/kafka/producer"
	"repository-hotel-booking/internal/app/router"
	"repository-hotel-booking/internal/app/service"
)

type Consumer struct {
	ConsumerConn *kafka.Conn
}

func (k *Consumer) ReadMessage(s *service.Service, producer *producer.Producer) {
	log.Println("Receiving messages")
	for {
		m, err := k.ConsumerConn.ReadMessage(10e6)
		if err != nil {
			break
		}
		req := &KafkaRequest{}
		json.Unmarshal(m.Value, req)
		router.DeliveryService(req.ServiceName, &req.Payload, producer, s)
		fmt.Printf("Received message %+v", req)
	}
	if err := k.ConsumerConn.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

type KafkaRequest struct {
	ServiceName string      `json:"service_name"`
	Payload     interface{} `json:"payload"`
}
