package model

type UserAddRequest struct {
	Id          string `json:"id"`
	Identifier  string `json:"identifier"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	ReferenceId string `json:"reference_id"`
}
type UserUpdateRequest struct {
	Password    string `json:"password"`
	Role        string `json:"role"`
	ReferenceId string `json:"reference_id"`
}
