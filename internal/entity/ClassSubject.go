package entity

import "time"

type ClassSubject struct {
	Id        int       `json:"id"`
	ClassId   int       `json:"class_id"`
	SubjectId int       `json:"subject_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Class     *Class    `json:"class"`
	Subject   *Subject  `json:"subject"`
}
