package model

import "time"

type Department struct {
	Id        int32
	Name      string
	ParentId  int32
	Status    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
