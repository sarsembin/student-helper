package usersvc

import "time"

type JsonEntity struct {
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	IsSuperuser bool      `json:"is_superuser"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	DateJoined  time.Time `json:"date_joined"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
}
