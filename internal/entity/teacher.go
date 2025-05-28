package entity

import "time"

type Teacher struct {
	Id        string    `json:"id"`
	Nip       string    `json:"nip"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Subjects  []Subject `json:"subjects"`
}
