package model

import "time"

type Role struct {
	Id        int32
	Name      string
	Title     string
	Status    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
