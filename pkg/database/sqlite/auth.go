package sqlite

import (
	"database/sql"
	"errors"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
)

type Auth struct {
	db *sql.DB
}

// Login verifies whether username exists and whether the password matches.
// It returns the id of the user if the authentication is successful.
// Otherwise, it returns an authentication error.
func (a *Auth) Login(username string, password string) (int, error) {
	var user models.User

	row := a.db.QueryRow(SELECT_USER_BY_FOR_AUTH, username)
	err := row.Scan(&user.Id, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, database.ErrAuth
		}

		return 0, errors.Join(database.ErrInternalDatabase, err)
	}

	if user.Password != password {
		return 0, database.ErrAuth
	}

	return user.Id, nil
}
