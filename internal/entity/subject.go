package entity

import "time"

type Subject struct {
	Id        int        `json:"id"`
	TeacherId string     `json:"teacher_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Teacher   *Teacher   `json:"teacher"`
	Schedules []Schedule `json:"schedules"`
	Grades    []Grade    `json:"grades"`
}
