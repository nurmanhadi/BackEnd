package entity

import "time"

type Class struct {
	Id        int        `json:"id"`
	TeacherId string     `json:"teacher_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Teacher   *Teacher   `json:"teacher"`
	Students  []Student  `json:"students"`
	Schedules []Schedule `json:"schedules"`
}
