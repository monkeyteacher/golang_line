package controllers

import (
	"golang_line/app/models/responses"
	"golang_line/app/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func UserController() *userController {
	return &userController{}
}

func (h *userController) GetUserMessage() gin.HandlerFunc {
	return func(context *gin.Context) {
		queryParams := context.Request.URL.Query()
		if len(queryParams) == 0 {
			message, err := services.UserService().GetALLMessage()
			if err != nil {
				context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
			} else {
				context.JSON(http.StatusOK, responses.Status(responses.Success, message))
			}
			return
		} else if len(queryParams) == 1 {
			if _, ok := queryParams["userID"]; ok {
				userID := context.Query("userID")
				if _, err := services.UserService().CheckUserExist(userID); err != nil {
					log.Println(err)
					context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
					return
				}
				message, err := services.UserService().GetUserMessagebyUserID(userID)
				if err != nil {
					context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
				} else {
					context.JSON(http.StatusOK, responses.Status(responses.Success, message))
				}
				return
			} else {
				context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
			}
		} else {
			context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
			return
		}
	}
}
