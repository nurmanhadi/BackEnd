package entity

import "time"

type Teacher struct {
	Id        string    `json:"id"`
	Nip       string    `json:"nip"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Class     *Class    `json:"class"`
}
