package universitescoredb

//nolint:go-staticcheck //def
type Entity struct {
	tableName struct{} `pg:"studentHelper_universitescore" `
	ID        int      `pg:"id,pk"`
	Title     string   `pg:"title"`
	Eval      string   `pg:"eval"`
}
