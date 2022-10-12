package customer_order

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOrderByParam(db *mongo.Database) func(context.Context, primitive.ObjectID, Order) error {
	return func(ctx context.Context, objectId primitive.ObjectID, order Order) error {
		collection := getCustomerCollection(db)
		filter := bson.M{"_id": objectId}
		res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": order})

		if res.ModifiedCount == 0 {
			return errors.New("Order updated")
		}
		return err
	}
}

func UpdateOrderByField(db *mongo.Database) func(context.Context, Order) error {
	return func(ctx context.Context, order Order) error {
		collection := getCustomerCollection(db)
		filter := bson.M{"_id": order.ID}
		res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": order})

		if res.ModifiedCount == 0 {
			return errors.New("order can not update")
		}
		return err
	}
}
