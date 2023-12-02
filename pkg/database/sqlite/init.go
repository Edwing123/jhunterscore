package sqlite

import (
	_ "embed"
	"errors"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"github.com/mattn/go-sqlite3"
)

//go:embed data/base.db
var INITIALIZED_DATABASE_DATA []byte

func isUniqueViolationErr(err sqlite3.Error) bool {
	return err.ExtendedCode == sqlite3.ErrConstraintUnique
}

func internalDatabaseErr(err error) error {
	return errors.Join(database.ErrInternalDatabase, err)
}
