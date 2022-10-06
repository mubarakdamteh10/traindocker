package users

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func getUserCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("users")
}
