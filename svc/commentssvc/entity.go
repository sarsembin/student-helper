package commentssvc

type JsonEntity struct {
	Content      string `json:"content"`
	UserID       int    `json:"user_id"`
	UniversityID int    `json:"university_id"`
}
