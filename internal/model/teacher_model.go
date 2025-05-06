package model

type TeacherUpdateRequest struct {
	Nip  string `json:"nip" validate:"omitempty,max=100"`
	Name string `json:"name" validate:"omitempty,max=100"`
}
type TeacherAddRequest struct {
	Nip      string `json:"nip" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
type TeacherResponse struct {
	Id   string `json:"id"`
	Nip  string `json:"nip"`
	Name string `json:"name"`
}
