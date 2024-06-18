package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Verifikasi struct {
	Username  string             `json:"username" bson:"username"`
	Kode      string             `json:"kode" bson:"kode"`
	NoWa      string             `json:"nowa" bson:"nowa"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
