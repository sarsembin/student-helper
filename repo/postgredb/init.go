package postgredb

import (
	"studentHelper/repo/postgredb/universitescoredb"
	"studentHelper/repo/postgredb/universitydb"

	"github.com/go-pg/pg/v10"
)

type Repository struct {
	Universitydb      universitydb.Repository
	Universitescoredb universitescoredb.Repository
}

func New(pg *pg.DB) *Repository {
	universityDB := universitydb.New(pg)
	Universitescoredb := universitescoredb.New(pg)

	return &Repository{
		Universitydb: universityDB,
		Universitescoredb: Universitescoredb,
	}
}
