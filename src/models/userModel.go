package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Age      int32              `bson:"age"`
	CreateAt time.Time          `bson:"createdAt"`
}
