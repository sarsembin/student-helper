package commentsdb

//nolint:unused //def
type Entity struct {
	tableName    struct{} `pg:"studentHelper_comments"`
	ID           int      `pg:"id,pk"`
	Content      string   `pg:"content"`
	UserID       int      `pg:"user_id_id,fk:studentHelper_user_id"`
	UniversityID int      `pg:"university_id_id,fk:studentHelper_university_id"`
}
