package res

import (
	"studentHelper/migrations"

	"github.com/go-pg/pg/v10"
)

type Resources struct {
	Postgres *pg.DB
}

func Init() (*Resources, error) {
	pg, err := newPostgres()
	if err != nil {
		return nil, err
	}

	err = migrations.Migrate(pg)
	if err != nil {
		return nil, err
	}

	return &Resources{
		Postgres: pg,
	}, nil
}
