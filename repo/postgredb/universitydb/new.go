package universitydb

import "github.com/go-pg/pg/v10"

type Repository interface {
	Add(e *University) error
}

type repository struct {
	db *pg.DB
}

func New(db *pg.DB) Repository {
	return &repository{
		db: db,
	}
}

func (c *repository) Add(e *University) error {
	_, err := pg.Model(e).Insert()
	return err
}
