package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	_id         primitive.ObjectID
	UserID      string `bson:"user_id" binding:"required"`
	Message     string `bson:"message" binding:"required"`
	CreatedTime time.Time
}
