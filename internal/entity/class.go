package entity

import "time"

type Class struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Students     []Student      `json:"students"`
	ClassSubject []ClassSubject `json:"class_subjects"`
}
