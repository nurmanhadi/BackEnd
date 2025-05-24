package model

type StudentAddRequest struct {
	Nis      string `json:"nis" validate:"required,max=100"`
	ClassId  int    `json:"class_id"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Status   string `json:"status" validate:"required,max=100"`
}
type StudentUpdateRequest struct {
	Nis    string `json:"nis" validate:"omitempty,max=100"`
	Name   string `json:"name" validate:"omitempty,max=100"`
	Status string `json:"status" validate:"omitempty,max=100"`
}
type StudentLoginRequest struct {
	Nis      string `json:"nis" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
type StudentResponse struct {
	Id      string `json:"id"`
	Nis     string `json:"nis"`
	ClassId int    `json:"class_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}
