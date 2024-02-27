package service

import (
	"encoding/json"
	"fmt"
	"log"
	"repository-hotel-booking/internal/app/kafka/producer"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/repository"
)

type Service struct {
	repo *repository.Repositories
}

func NewService(repo *repository.Repositories) *Service {
	return &Service{repo: repo}
}
func (s *Service) DeliveryService(serviceName string, payload interface{}, producer *producer.Producer) {
	switch serviceName {
	case "GetAccounts":
		jsonbody, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err)
		}
		var query *model.AccountQuery
		if err := json.Unmarshal(jsonbody, &query); err != nil {
			// do error check
			fmt.Println(err)
		}
		log.Println(query)
		accounts, errI := s.GetAccounts(query)
		response := &model.Response{
			Error: errI,
			Body:  accounts,
		}
		log.Println(accounts)

		message, err := json.Marshal(response)
		log.Println("Sent to broker", response)
		producer.SendMessage(message)

	}
}
