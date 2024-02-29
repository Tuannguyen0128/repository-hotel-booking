package controller

import (
	"encoding/json"
	"fmt"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/service"
)

func GetStaffs(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var query *model.StaffQuery
	if err = json.Unmarshal(jsonBody, &query); err != nil {
		fmt.Println(err)
	}

	staffs, errI := service.GetStaffs(query)
	response := &model.Response{
		Error: errI,
		Body:  staffs,
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}

func CreateStaff(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var staff *model.Staff
	if err := json.Unmarshal(jsonBody, &staff); err != nil {
		// do error check
		fmt.Println(err)
	}
	id, errI := service.AddStaff(staff)
	response := &model.Response{
		Error: errI,
		Body:  model.AddStaffResponse{ID: id},
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}

func DeleteStaff(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	staff := &model.DeleteStaffRequest{}
	if err = json.Unmarshal(jsonBody, &staff); err != nil {
		// do error check
		fmt.Println(err)
	}
	result, errI := service.DeleteStaff(staff.ID)
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

func UpdateStaff(payload *interface{}, service *service.Service) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	var staff *model.Staff
	if err := json.Unmarshal(jsonBody, &staff); err != nil {
		// do error check
		fmt.Println(err)
	}
	staff, errI := service.UpdateStaff(staff)
	response := &model.Response{
		Error: errI,
		Body:  staff,
	}

	message, err := json.Marshal(response)
	if err != nil {
		messageNil, _ := json.Marshal(response)
		return messageNil, err
	}
	return message, nil
}
