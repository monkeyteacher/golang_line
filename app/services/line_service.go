package services

import (
	"golang_line/app/repositories"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineService struct {
}

func LineService() *lineService {
	return &lineService{}
}

func (h *lineService) LineMessageHandler(MyBot *linebot.Client, events []*linebot.Event) {
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "查詢UserID" {
					if _, err := MyBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.Source.UserID)).Do(); err != nil {
						log.Println(err)
					}
				} else {
					repositories.MessageRepository().Create(event.Source.UserID, message.Text)
					if _, err := MyBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("收到")).Do(); err != nil {
						log.Println(err)
					}

				}
			}
		}
	}
}
