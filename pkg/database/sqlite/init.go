package sqlite

import (
	"database/sql"
	_ "embed"
	"errors"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
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

type Database struct {
	Users     models.UsersRepository
	Auth      models.AuthRepository
	Files     models.FilesRepository
	Companies models.CompaniesRepository
}

func New(db *sql.DB) models.Database {
	return models.Database{
		AuthRepository:      &Auth{db},
		UsersRepository:     &Users{db},
		FilesRepository:     &Files{db},
		CompaniesRepository: &Companies{db},
	}
}
