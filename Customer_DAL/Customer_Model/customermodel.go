package customermodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Bank_id   string  `json:"bank_id" bson:"bank_id"`
	Balance   float64 `json:"balance" bson:"balance"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	IsActive  bool    `json:"is_active" bson:"is_active"`
}
type Tranaction struct {
	From_coustomer string    `json:"from" bson:"from"`
	To_coustomer   string    `json:"to" bson:"to"`
	Amount         float64   `json:"amount" bson:"amount"`
}
type TranactionResponse struct {
	From_coustomer string    `json:"from" bson:"from"`
	To_coustomer   string    `json:"to" bson:"to"`
	Amount         float64   `json:"amount" bson:"amount"`
}
type CustomerResponse struct {
	Customer_Id primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
}
