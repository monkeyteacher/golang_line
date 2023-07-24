package repositories

import (
	"context"
	"golang_line/app/models"
	"golang_line/database"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (h *messageRepository) GetMessagesbyUserID(userID string) ([]models.Message, error) {
	return h.findOnebyQuery(bson.M{"user_id": userID})
}

func (h *messageRepository) GetAllMessages() ([]models.Message, error) {
	return h.findOnebyQuery(bson.M{})
}

func (h *messageRepository) findOnebyQuery(query interface{}) ([]models.Message, error) {
	messageCollection := database.MI.DB.Collection("message")
	var messages []models.Message
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := messageCollection.Find(ctx, query)
	if err == mongo.ErrNoDocuments {
		log.Fatal(err)
	}

	for results.Next(ctx) {
		var singleMessage models.Message
		if err = results.Decode(&singleMessage); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, singleMessage)
	}
	return messages, err
}
