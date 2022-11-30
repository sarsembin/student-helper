package res

import (
	"github.com/go-pg/pg/v10"
)

type Resources struct {
	Postgres *pg.DB
}

func New() (*Resources, error) {
	pg, err := newPostgres()
	if err != nil {
		return nil, err
	}

	return &Resources{
		Postgres: pg,
	}, nil
}
