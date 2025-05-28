package entity

import "time"

type Schedule struct {
	Id         int64        `json:"id"`
	SubjectId  int          `json:"subject_id"`
	Timetable  time.Time    `json:"timetable"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Subject    *Subject     `json:"subject"`
	Attendaces []Attendance `json:"attendaces"`
}
