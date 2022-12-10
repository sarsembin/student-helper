package userdb

import "time"

//nolint:unused //def
type Entity struct {
	tableName   struct{}  `pg:"studentHelper_user"`
	ID          int       `pg:"id,pk"`
	Password    string    `pg:"password,notnull"`
	LastLogin   time.Time `pg:"last_login"`
	IsSuperuser bool      `pg:"is_superuser,use_zero"`
	FirstName   string    `pg:"first_name,unique"`
	LastName    string    `pg:"last_name"`
	IsStaff     bool      `pg:"is_staff,use_zero"`
	IsActive    bool      `pg:"is_active"`
	DateJoined  time.Time `pg:"date_joined"`
	Username    string    `pg:"username,unique,notnull"`
	Email       string    `pg:"email,unique,notnull"`
}
