package util

import "repository-hotel-booking/internal/app/model"

func BuildResponse(errInfo *model.ErrInfo, body interface{}) model.Response {
	return model.Response{
		Error: errInfo,
		Body:  body,
	}
}
