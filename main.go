package main

import (
	"log"
	"repository-hotel-booking/internal/app"
	"repository-hotel-booking/internal/app/util"
)

func main() {
	const configPath = "./configs"
	if err := util.LoadConfigs("", configPath, "app", app.CFG); err != nil {
		log.Panicf("error while reading app configs %v", err)
	}

	//mq := kafka.InitConnection()
	//mq.ConsumeMessage()
	app.Init()
}
