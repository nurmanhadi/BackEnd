package model

type AuthResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
type AuthLoginRequest struct {
	Identifier string `json:"identifier" validate:"required,max=100"`
	Password   string `json:"password" validate:"required,max=100"`
}
