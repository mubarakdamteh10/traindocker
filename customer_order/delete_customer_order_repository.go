package customer_order

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOrderById(db *mongo.Database) func(context.Context, primitive.ObjectID) error {
	return func(ctx context.Context, objectId primitive.ObjectID) error {
		collection := getCustomerCollection(db)
		filter := bson.M{"_id": objectId}
		res, err := collection.DeleteOne(ctx, filter)
		if res.DeletedCount == 0 {
			return errors.New("order deleted")
		}
		return err
	}
}
