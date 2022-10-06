package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"Name" json:"Name"`
	Location string             `bson:"Location" json:"Location"`
	Title    string             `bson:"Title" json:"Title"`
}
