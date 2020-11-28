package model

import "time"

type Role struct {
	ID        string     `json:"id"`
	Data      string     `json:"data"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Rolee struct {
	ID        string     `json:"id"`
	Data      Rollee     `json:"data"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Rollee struct {
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}
