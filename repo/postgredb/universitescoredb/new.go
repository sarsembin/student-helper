package universitescoredb

import "github.com/go-pg/pg/v10"

type Repository interface {
	Add(e *Entity) (err error)
	GetAll() (res []Entity, err error)
	Get(id int) (res *Entity, err error)
	Put(e *Entity) (err error)
	Delete(id int) (err error)
}

type repository struct {
	db *pg.DB
}

func New(db *pg.DB) Repository {
	return &repository{
		db: db,
	}
}
