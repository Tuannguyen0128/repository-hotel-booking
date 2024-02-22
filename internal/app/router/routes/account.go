package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/service"
	"strconv"
)

func AccountRoutes(s *service.Service) []Route {
	return []Route{
		{
			Uri:    "/accounts",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				page, _ := strconv.Atoi(c.Query("page"))
				size, _ := strconv.Atoi(c.Query("size"))
				accountQuery := &model.AccountQuery{
					ID:       c.Query("id"),
					StaffID:  c.Query("staff_id"),
					Username: c.Query("username"),
					Page:     page,
					Size:     size,
				}
				accounts, err := s.GetAccounts(accountQuery)
				if err != nil {
					fmt.Println(err.Error())
					c.JSON(http.StatusBadRequest, nil)
					return
				}
				c.JSON(http.StatusOK, model.Accounts{Accounts: accounts})
			},
		},
		//{
		//	Uri:          "/account/{id}",
		//	Method:       http.MethodGet,
		//	Handler:      service.GetAccount,
		//},
		//{
		//	Uri:          "/account",
		//	Method:       http.MethodPost,
		//	Handler:      service.CreateAccount,
		//},
		//{
		//	Uri:          "/account/{id}",
		//	Method:       http.MethodPut,
		//	Handler:      service.UpdateAccount,
		//},
		//{
		//	Uri:          "/account/{id}",
		//	Method:       http.MethodDelete,
		//	Handler:      service.DeleteAccount,
		//},
		//{
		//	Uri:          "/get-account-by-email",
		//	Method:       http.MethodGet,
		//	Handler:      service.FindByEmail,
		//},
		//{
		//	Uri:          "/get-account-by-merchantcode",
		//	Method:       http.MethodGet,
		//	Handler:      service.FindByMerchantCode,
		//},
	}
}
