package model

type ScheduleAddRequest struct {
	ClassId   int    `json:"class_id" validate:"required"`
	SubjectId int    `json:"subject_id" validate:"required"`
	Timetable string `json:"timetable" validate:"required"`
}
type ScheduleUpdateRequest struct {
	ClassId   int    `json:"class_id" validate:"omitempty"`
	SubjectId int    `json:"subject_id" validate:"omitempty"`
	Timetable string `json:"timetable" validate:"omitempty"`
}
