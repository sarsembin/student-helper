package postgredb

import (
	"studentHelper/repo/postgredb/universitydb"

	"github.com/go-pg/pg/v10"
)

type Repository struct {
	Universitydb universitydb.Repository
}

func New(pg *pg.DB) *Repository {
	universityDB := universitydb.New(pg)

	return &Repository{
		Universitydb: universityDB,
	}
}
