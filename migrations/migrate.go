package migrations

import (
	"fmt"
	"io/fs"
	"os"
	"sort"

	"strconv"
	"strings"

	"github.com/go-pg/pg/v10"
)

const migrationVersion = 10

// folder name in the project folder
const dirName = "./migrations/"

type UpDown int

const (
	TypeUnknown = iota
	TypeUp
	TypeDown
)

type SqlMigration struct {
	Version  int
	UpDown   UpDown
	DirEntry fs.DirEntry
}

func Migrate(pg *pg.DB) error {
	dir, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	migraions, err := extractType(dir)
	if err != nil {
		return err
	}

	sort.Slice(migraions, func(i, j int) bool {
		return migraions[i].Version < migraions[j].Version
	})

	// TODO: refactor this poop
	for i := migraions[len(migraions)-1].Version; i > migrationVersion; i-- {
		if migraions[i].UpDown == TypeDown {
			content, err := os.ReadFile(dirName + migraions[i].DirEntry.Name())
			if err != nil {
				return err
			}

			_, err = pg.Exec(string(content))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	for i := 0; i <= migrationVersion; i++ {
		if migraions[i].UpDown == TypeUp {
			content, err := os.ReadFile(dirName + migraions[i].DirEntry.Name())
			if err != nil {
				return err
			}

			_, err = pg.Exec(string(content))
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

// TODO: fix for edge cases
func extractType(dirEntry []fs.DirEntry) (migrations []SqlMigration, err error) {
	for _, f := range dirEntry {
		split := strings.Split(f.Name(), "_")
		if len(split) == 1 {
			continue
		}

		version, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}

		switch {
		case strings.HasSuffix(f.Name(), ".up.sql") && version <= migrationVersion:
			migrations = append(migrations, SqlMigration{
				Version:  version,
				UpDown:   TypeUp,
				DirEntry: f,
			})
		case strings.HasSuffix(f.Name(), ".down.sql") && version > migrationVersion:
			migrations = append(migrations, SqlMigration{
				Version:  version,
				UpDown:   TypeDown,
				DirEntry: f,
			})
		}
	}

	return migrations, nil
}
