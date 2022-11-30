package res

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	pass = "1"
	db   = "back"
)

func newPostgres() (con *pg.DB, err error) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")

	options := &pg.Options{
		User:     user,
		Password: pass,
		Addr:     address,
		Database: db,
		PoolSize: 50,
	}

	con = pg.Connect(options)
	if con == nil {
		return nil, fmt.Errorf("cannot connect to postgres")
	}

	return con, nil
}
