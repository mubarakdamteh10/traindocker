package users

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(db *mongo.Database) func(context.Context, *Users) error {
	return func(ctx context.Context, user *Users) error {
		collection := getUserCollection(db)
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
		return nil
	}

}
