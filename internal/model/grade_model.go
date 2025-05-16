package model

type GradeAddRequest struct {
	StudentId       string `json:"student_id" validate:"required,max=36"`
	SubjectId       int    `json:"subject_id" validate:"required"`
	AttendanceScore int    `json:"attendance_score" validate:"omitempty"`
	TaskScore       int    `json:"task_score" validate:"omitempty"`
	MidternScore    int    `json:"midtern_score" validate:"omitempty"`
	FinalScore      int    `json:"final_score" validate:"omitempty"`
}
type GradeUpdateRequest struct {
	StudentId       string `json:"student_id" validate:"omitempty,max=36"`
	SubjectId       int    `json:"subject_id" validate:"omitempty"`
	AttendacenScore int    `json:"attendance_score" validate:"omitempty"`
	TaskScore       int    `json:"task_score" validate:"omitempty"`
	MidternScore    int    `json:"midtern_score" validate:"omitempty"`
	FinalScore      int    `json:"final_score" validate:"omitempty"`
}
