package repo

import (
	"studentHelper/repo/postgredb"
	"studentHelper/res"
)

type Repositories struct {
	Postgrerepo postgredb.Repository
}

func New(res res.Resources) *Repositories {
	postgredb.New()

	return &Repositories{
		res: res,
	}
}
