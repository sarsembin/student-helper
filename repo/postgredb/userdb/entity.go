package userdb

//nolint:unused //def
type Entity struct {
	tableName   struct{} `pg:"studentHelper_user"`
	ID          int      `pg:"id,pk"`
	Password    string   `pg:"password,notnull"`
	LastLogin   string   `pg:"last_login"`
	IsSuperuser string   `pg:"is_superuser"`
	FirstName   string   `pg:"first_name,unique"`
	LastName    string   `pg:"last_name"`
	IsStaff     string   `pg:"is_staff"`
	IsActive    string   `pg:"is_active"`
	DateJoined  string   `pg:"date_joined"`
	Username    string   `pg:"username,unique,notnull"`
	Email       string   `pg:"email,unique,notnull"`
}
