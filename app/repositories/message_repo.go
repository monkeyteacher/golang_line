package repositories

import (
	"context"
	"golang_line/app/models"
	"golang_line/database"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type messageRepository struct {
}

func MessageRepository() *messageRepository {
	return &messageRepository{}
}

func (h *messageRepository) Create(userID string, lineMessage string) *mongo.InsertOneResult {
	messageCollection := database.MI.DB.Collection("message")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	message := models.Message{
		UserID:      userID,
		Message:     lineMessage,
		CreatedTime: time.Now(),
	}
	result, err := messageCollection.InsertOne(ctx, message)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
