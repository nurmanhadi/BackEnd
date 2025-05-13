package model

type AttendanceAddRequest struct {
	StudentId  string `json:"student_id" validate:"required,max=36"`
	ScheduleId int64  `json:"schedule_id" validate:"required"`
	Status     string `json:"status" validate:"required,max=20"`
}
type AttendanceUpdateRequest struct {
	StudentId  string `json:"student_id" validate:"omitempty,max=36"`
	ScheduleId int64  `json:"schedule_id" validate:"omitempty"`
	Status     string `json:"status" validate:"omitempty,max=20"`
}
