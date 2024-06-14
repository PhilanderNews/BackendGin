package models

type CredentialUser struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    struct {
		Name     string `json:"name" bson:"name"`
		Username string `json:"username" bson:"username"`
		Role     string `json:"role" bson:"role"`
	} `json:"data" bson:"data"`
}

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
	Empty   bool        `json:"empty,omitempty" bson:"empty,omitempty"`
}
