package main

import (
	"log"
	"repository-hotel-booking/configs"
	"repository-hotel-booking/internal/app"
	"repository-hotel-booking/internal/app/kafka"
	"repository-hotel-booking/internal/app/kafka/consumer"
	"repository-hotel-booking/internal/app/kafka/producer"
	"repository-hotel-booking/internal/app/repository"
	"repository-hotel-booking/internal/app/service"
	"repository-hotel-booking/internal/app/util"
)

var KafkaConsumer *consumer.Consumer
var KafkaProducer *producer.Producer

func main() {
	const configPath = "./configs"
	if err := configs.LoadConfigs("", configPath, "app", app.CFG); err != nil {
		log.Panicf("error while reading app configs %v", err)
	}
	db := util.InitConnection(app.CFG.DB)
	repo := repository.New(db)
	KafkaConsumer, KafkaProducer = kafka.InitConnection(app.CFG.KafkaServer, app.CFG.KafkaRepoTopic, app.CFG.KafkaBrokerTopic, 0)
	s := service.NewService(repo)
	KafkaConsumer.ReadMessage(s, KafkaProducer)
	//app.Init()
}
