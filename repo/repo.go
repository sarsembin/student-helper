package repo

import (
	"studentHelper/repo/postgredb"
	"studentHelper/res"
)

type Repositories struct {
	Postgrerepo postgredb.Repository
}

func Init(res *res.Resources) *Repositories {
	postgredb := postgredb.New(res.Postgres)

	return &Repositories{
		Postgrerepo: *postgredb,
	}
}
