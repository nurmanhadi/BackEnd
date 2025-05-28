package model

type ClassAddRequest struct {
	Name string `jason:"name" validate:"required,max=100"`
}
type ClassUpdateRequest struct {
	Name string `jason:"name" validate:"omitempty,max=100"`
}
