package model

type ClassSubjectAddRequest struct {
	ClassId   int `json:"class_id" validate:"required"`
	SubjectId int `json:"subject_id" validate:"required"`
}
