package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type IngredientDto struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"ingredientName" bson:"ingredientName,omitempty"`
	Quantity int                `json:"quantity" bson:"quantity,omitempty"`
}
