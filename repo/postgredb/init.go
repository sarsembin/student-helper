package postgredb

import (
	"studentHelper/repo/postgredb/commentsdb"
	"studentHelper/repo/postgredb/universitescoredb"
	"studentHelper/repo/postgredb/universitydb"
	"studentHelper/repo/postgredb/useracademicinfodb"
	"studentHelper/repo/postgredb/userdb"

	"github.com/go-pg/pg/v10"
)

type Repository struct {
	Universitydb       universitydb.Repository
	Universitescoredb  universitescoredb.Repository
	Userdb             userdb.Repository
	UserAcademicInfodb useracademicinfodb.Repository
	Commentsdb         commentsdb.Repository
}

func New(pg *pg.DB) *Repository {
	universityDB := universitydb.New(pg)
	Universitescoredb := universitescoredb.New(pg)
	Userdb := userdb.New(pg)
	UserAcademicInfodb := useracademicinfodb.New(pg)
	Commentsdb := commentsdb.New(pg)

	return &Repository{
		Universitydb:       universityDB,
		Universitescoredb:  Universitescoredb,
		Userdb:             Userdb,
		UserAcademicInfodb: UserAcademicInfodb,
		Commentsdb:         Commentsdb,
	}
}
