package universitescoresvc

type JsonEntity struct {
	Title string `json:"title"`
	Eval  string `json:"eval"`
	UniID int    `json:"university_id"`
}
