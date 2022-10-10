package customer_order

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func getCustomerCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("customer_order")
}
