package usersvc

type JsonEntity struct {
	Password    string `json:"password"`
	LastLogin   string `json:"last_login"`
	IsSuperuser string `json:"is_superuser"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsStaff     string `json:"is_staff"`
	IsActive    string `json:"is_active"`
	DateJoined  string `json:"date_joined"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}
