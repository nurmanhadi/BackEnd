package entity

import "time"

type User struct {
	Id          string    `json:"id"`
	Identifier  string    `json:"identifier"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	ReferenceId string    `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
