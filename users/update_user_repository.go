package users

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUserByIdParam(db *mongo.Database) func(context.Context, primitive.ObjectID, Users) error {
	return func(ctx context.Context, objectId primitive.ObjectID, user Users) error {
		collection := getUserCollection(db)
		filter := bson.M{"_id": objectId}
		res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": user})

		if res.ModifiedCount == 0 {
			return errors.New("user Updated")
		}
		return err
	}
}

func UpdateUserByIdField(db *mongo.Database) func(context.Context, Users) error {
	return func(ctx context.Context, user Users) error {
		collection := getUserCollection(db)
		filter := bson.M{"_id": user.ID}
		res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": user.ID})

		if res.ModifiedCount == 0 {
			return errors.New("users caon not update")
		}
		return err
	}
}
