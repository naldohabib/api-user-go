package model

type UserList struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}
