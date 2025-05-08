package model

type SubjectAddRequest struct {
	TeacherId string `json:"teacher_id" validate:"required,max=36"`
	Name      string `json:"name" validate:"required,max=100"`
}
type SubjectUpdateRequest struct {
	TeacherId string `json:"teacher_id" validate:"omitempty,max=36"`
	Name      string `json:"name" validate:"omitempty,max=100"`
}
