package users

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUserById(db *mongo.Database) func(context.Context, primitive.ObjectID) error {
	return func(ctx context.Context, objectId primitive.ObjectID) error {
		colletion := getUserCollection(db)
		filter := bson.M{"_id": objectId}
		res, err := colletion.DeleteOne(ctx, filter)
		if res.DeletedCount == 0 {
			return errors.New("users can not delete")
		}
		return err
	}
}
