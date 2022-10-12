package customer_order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string               `bson:"Name" json:"name"`
	Amount  primitive.Decimal128 `bson:"amount" json:"amount"`
	Address string               `bson:"address" json:"address"`
	Payment bool                 `bson:"payment" json:"payment"`
}
