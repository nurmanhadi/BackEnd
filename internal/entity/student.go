package entity

import "time"

type Student struct {
	Id         string       `json:"id"`
	Nis        string       `json:"nis"`
	ClassId    int          `json:"class_id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Password   string       `json:"-"`
	Status     string       `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Class      *Class       `json:"class"`
	Attendaces []Attendance `json:"attendaces"`
}
