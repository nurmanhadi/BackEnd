package entity

import "time"

type Attendance struct {
	Id         int64     `json:"id"`
	StudentId  string    `json:"student_id"`
	ScheduleId int64     `json:"schedule_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Student    *Student  `json:"student"`
	Schedule   *Schedule `json:"schedule"`
}
