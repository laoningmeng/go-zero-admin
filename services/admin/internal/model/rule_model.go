package model

import "time"

type Rule struct {
	Id        int64
	Name      string
	Title     string
	Type      int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
