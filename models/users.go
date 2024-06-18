package models

type Users struct {
	Name     string `json:"name" bson:"name"`
	NoWa     string `json:"nowa" bson:"nowa"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Kode     string `json:"kode" bson:"kode"`
	Role     string `json:"role" bson:"role"`
}
