package controllers

import (
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
