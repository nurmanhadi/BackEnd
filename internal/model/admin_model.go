package model

type AdminAddRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"Password" validate:"required,max=100"`
}
type AdminUpdateRequest struct {
	Username string `json:"username" validate:"omitempty,max=100"`
}
