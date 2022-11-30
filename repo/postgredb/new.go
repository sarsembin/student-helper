package postgredb

import "github.com/go-pg/pg/v10"

type Repository interface {
}

type repository struct {
	conn *pg.DB
}

func New(db *pg.DB) Repository {
	return &repository{
		conn: db,
	}
}
