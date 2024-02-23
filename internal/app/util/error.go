package util

import "repository-hotel-booking/internal/app/model"

func BuildErrInfo(code, msg string) *model.ErrInfo {
	return &model.ErrInfo{
		Code:    code,
		Message: msg,
	}
}
