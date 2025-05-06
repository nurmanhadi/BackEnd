package model

type StudentRegisterRequest struct {
	Nis      string `json:"nis" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Status   string `json:"status" validate:"required,max=100"`
}
