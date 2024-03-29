package router

import (
	"fmt"
	"repository-hotel-booking/internal/app/controller"
	"repository-hotel-booking/internal/app/kafka/producer"
	"repository-hotel-booking/internal/app/service"
)

func DeliveryService(serviceName string, payload *interface{}, producer *producer.Producer, service *service.Service) {
	var message []byte
	var err error
	switch serviceName {
	case "GetAccounts":
		message, err = controller.GetAccounts(payload, service)
		break
	case "CreateAccount":
		message, err = controller.CreateAccount(payload, service)
		break
	case "DeleteAccount":
		message, err = controller.DeleteAccount(payload, service)
		break
	case "UpdateAccount":
		message, err = controller.UpdateAccount(payload, service)
		break
	case "GetStaffs":
		message, err = controller.GetStaffs(payload, service)
		break
	case "CreateStaff":
		message, err = controller.CreateStaff(payload, service)
		break
	case "DeleteStaff":
		message, err = controller.DeleteStaff(payload, service)
		break
	case "UpdateStaff":
		message, err = controller.UpdateStaff(payload, service)
		break
	}

	if err != nil {
		fmt.Println(err)
	}
	producer.SendMessage(message)
}
