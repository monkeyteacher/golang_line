package controllers

import (
	"golang_line/app/models"
	"golang_line/app/models/responses"
	"golang_line/app/services"
	"golang_line/configs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineController struct {
}

func LineController() *lineController {
	return &lineController{}
}

func (h *lineController) LineCallBack() gin.HandlerFunc {
	return func(context *gin.Context) {
		MyBot, err := linebot.New(
			configs.EnvConfigs.LineChannelSecret,
			configs.EnvConfigs.LineChannelToken,
		)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
			return
		}
		events, err := MyBot.ParseRequest(context.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
			} else {
				context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
			}
			return
		}
		services.LineService().LineMessageHandler(MyBot, events)
	}
}

func (h *lineController) SendMessage() gin.HandlerFunc {
	return func(context *gin.Context) {
		var userMessage models.UserMessage
		if err := context.ShouldBindJSON(&userMessage); err != nil {
			log.Println(err)
			context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
			return
		}

		MyBot, err := linebot.New(
			configs.EnvConfigs.LineChannelSecret,
			configs.EnvConfigs.LineChannelToken,
		)

		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
			return
		}

		if _, err := services.UserService().CheckUserExist(userMessage.UserID); err != nil {
			log.Println(err)
			context.JSON(http.StatusBadRequest, responses.Status(responses.ParameterErr, nil))
			return
		}

		err = services.LineService().SendMessageByUserID(MyBot, userMessage.UserID, userMessage.Message)
		if err != nil {
			context.JSON(http.StatusInternalServerError, responses.Status(responses.Error, nil))
		} else {
			context.JSON(http.StatusOK, responses.Status(responses.Success, nil))
		}
	}
}
