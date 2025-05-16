package entity

import "time"

type Grade struct {
	Id              int64     `json:"id"`
	StudentId       string    `json:"student_id"`
	SubjectId       int       `json:"schedule_id"`
	AttendanceScore int       `json:"attendance_score"`
	TaskScore       int       `json:"task_score"`
	MidternScore    int       `json:"midtern_score"`
	FinalScore      int       `json:"final_score"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Student         *Student  `json:"student"`
	Subject         *Subject  `json:"subject"`
}
