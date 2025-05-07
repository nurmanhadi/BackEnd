package model

type ClassAddRequest struct {
	TeacherId string `json:"teacher_id" validate:"required,max=36"`
	Name      string `jason:"name" validate:"required,max=100"`
}
type ClassUpdateRequest struct {
	TeacherId string `json:"teacher_id" validate:"omitempty,max=36"`
	ClassId   string `json:"class_id" validate:"omitempty"`
	Name      string `jason:"name" validate:"omitempty,max=100"`
}
