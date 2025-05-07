package entity

import "time"

type StudentClass struct {
	Id        int       `json:"id"`
	ClassId   int       `json:"class_id"`
	StudentId string    `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Student   *Student  `json:"student"`
	Class     Class     `json:"class"`
}
