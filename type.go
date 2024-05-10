package be63

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type solana struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Network   string             `bson:"network" json:"network"`
	Storage   string             `bson:"storage" json:"storage"`
	RAM       string             `bson:"ram" json:"ram"`
	Processor string             `bson:"processor" json:"processor"`
	Camera    int                `bson:"camera" json:"camera"`
	Display   int                `bson:"display" json:"display"`
	Battery   int                `bson:"battery" json:"battery"`
	OS        string             `bson:"os" json:"os"`
	Price     float64            `bson:"price" json:"price"`
}
