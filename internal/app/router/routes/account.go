package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/service"
	"repository-hotel-booking/internal/app/util"
	"strconv"
)

func AccountRoutes(s *service.Service) []Route {
	return []Route{
		// Get multiple account
		{
			Uri:    "/accounts",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				page, er := strconv.Atoi(c.Query("page"))

				size, er := strconv.Atoi(c.Query("size"))

				if er != nil {
					page = 0
					size = 0
				}

				accountQuery := &model.AccountQuery{
					ID:       c.Query("id"),
					StaffID:  c.Query("staff_id"),
					Username: c.Query("username"),
					Page:     page,
					Size:     size,
				}
				accounts, err := s.GetAccounts(accountQuery)
				if err.Code != "" {
					c.JSON(http.StatusBadRequest, util.BuildResponse(err, nil))
					return
				}
				c.JSON(http.StatusOK, util.BuildResponse(err,
					model.Accounts{
						Accounts: accounts,
					}))
			},
		},
		// Add one account
		{
			Uri:    "/account",
			Method: http.MethodPost,
			Handler: func(c *gin.Context) {
				newAccount := &model.Account{}
				c.ShouldBindJSON(newAccount)
				id, err := s.AddAccount(newAccount)
				if err.Code != "" {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, util.BuildResponse(err, nil))
					return
				}
				c.JSON(http.StatusOK, util.BuildResponse(err, model.AddAccountResponse{ID: id}))
			},
		},
		//Update one account
		{
			Uri:    "/account",
			Method: http.MethodPut,
			Handler: func(c *gin.Context) {
				newAccount := &model.Account{}
				c.ShouldBindJSON(newAccount)
				updatedAccount, err := s.UpdateAccount(newAccount)
				if err.Code != "" {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, util.BuildResponse(err, nil))
					return
				}
				c.JSON(http.StatusOK, util.BuildResponse(err, updatedAccount))
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
		{
			Uri:    "/account/:id",
			Method: http.MethodDelete,
			Handler: func(c *gin.Context) {
				id := c.Param("id")
				if id == "" {
					c.JSON(http.StatusBadRequest, util.BuildResponse(&model.ErrInfo{
						Code:    "E02",
						Message: "Invalid id",
					}, nil))
				}
				result, err := s.DeleteAccount(id)
				if err.Code != "" {
					log.Println(err.Error())
					c.JSON(http.StatusBadRequest, util.BuildResponse(err, nil))
					return
				}
				c.JSON(http.StatusOK, util.BuildResponse(err, result))
			},
		},
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
