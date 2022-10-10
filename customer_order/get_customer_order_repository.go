package customer_order

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllOrder(db *mongo.Database) func(context.Context) ([]*Order, error) {
	return func(ctx context.Context) ([]*Order, error) {
		collection := getCustomerCollection(db)
		var orders []*Order
		Cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		if err = Cursor.All(ctx, &orders); err != nil {
			return nil, err
		}
		fmt.Println(orders)
		return orders, nil
	}
}
