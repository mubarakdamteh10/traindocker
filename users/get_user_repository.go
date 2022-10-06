package users

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUser(db *mongo.Database) func(context.Context) ([]*Users, error) {
	return func(ctx context.Context) ([]*Users, error) {
		collection := getUserCollection(db)
		var users []*Users
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &users); err != nil {
			return nil, err
		}
		fmt.Println(users)
		return users, nil
	}
}

func GetUserById(db *mongo.Database) func(context.Context, string) (*Users, error) {
	return func(ctx context.Context, str string) (*Users, error) {
		collection := getUserCollection(db)
		filter := bson.M{"Name": str}
		var users Users
		if err := collection.FindOne(ctx, filter).Decode(&users); err != nil {
			return nil, err
		}
		fmt.Println(users)
		return &users, nil
	}
}
