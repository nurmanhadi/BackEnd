package entity

import "time"

type Schedule struct {
	Id        int64     `json:"id"`
	ClassId   int       `json:"class_id"`
	SubjectId int       `json:"subject_id"`
	Timetable time.Time `json:"timetable"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Class     *Class    `json:"class"`
	Subject   *Subject  `json:"subject"`
}
