package models

import "time"

type Payload struct {
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Exp      time.Time `json:"exp"`
	Iat      time.Time `json:"iat"`
	Nbf      time.Time `json:"nbf"`
}
