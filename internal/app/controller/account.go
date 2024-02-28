package controller

import (
	"encoding/json"
	"fmt"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/service"
)

func GetAccounts(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var query *model.AccountQuery
	if err = json.Unmarshal(jsonBody, &query); err != nil {
		fmt.Println(err)
	}

	accounts, errI := service.GetAccounts(query)
	response := &model.Response{
		Error: errI,
		Body:  accounts,
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}

func CreateAccount(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var account *model.Account
	if err := json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
	id, errI := service.AddAccount(account)
	response := &model.Response{
		Error: errI,
		Body:  model.AddAccountResponse{ID: id},
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}

func DeleteAccount(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	account := &model.DeleteAccountRequest{}
	if err = json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
	result, errI := service.DeleteAccount(account.ID)
	response := &model.Response{
		Error: errI,
		Body:  model.DeleteAccountResponse{Result: result},
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}

func UpdateAccount(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var account *model.Account
	if err := json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
	account, errI := service.UpdateAccount(account)
	response := &model.Response{
		Error: errI,
		Body:  account,
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}
