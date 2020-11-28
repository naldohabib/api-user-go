package model

import (
	"time"
)

type User struct {
	ID        string     `json:"id"`
	Data      string     `json:"data"`
	RoleId    string     `json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Users struct {
	ID        string     `json:"id"`
	Data      UserWrap   `json:"data"`
	RoleId    string     `json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type UserWrap struct {
	Email    string  `json:"email"`
	Status   Statuss `json:"status"`
	Password string  `json:"password"`
	Username string  `json:"username"`
}

type Statuss struct {
	Is_active bool `json:"is_active"`
}
