package model

type ScheduleAddRequest struct {
	SubjectId int    `json:"subject_id" validate:"required"`
	Timetable string `json:"timetable" validate:"required"`
}
type ScheduleUpdateRequest struct {
	SubjectId int    `json:"subject_id" validate:"omitempty"`
	Timetable string `json:"timetable" validate:"omitempty"`
}
