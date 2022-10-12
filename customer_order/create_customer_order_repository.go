package customer_order

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCustomerOrder(db *mongo.Database) func(context.Context, *Order) error {
	return func(ctx context.Context, order *Order) error {
		collection := getCustomerCollection(db)
		_, err := collection.InsertOne(ctx, order)
		if err != nil {
			return err
		}
		return nil
	}
}
