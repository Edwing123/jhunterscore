package sqlite

import (
	"database/sql"
	"errors"
	"strings"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/mattn/go-sqlite3"
)

type Users struct {
	db *sql.DB
}

func (u *Users) isUsernameTakenErr(msg string) bool {
	return strings.Contains(msg, "username")
}

func (u *Users) isFirstNameTakenErr(msg string) bool {
	return strings.Contains(msg, "first_name")
}

func (u *Users) getUserErr(msg string) (err error) {
	switch true {
	case u.isUsernameTakenErr(msg):
		err = database.ErrUsernameTaken
	case u.isFirstNameTakenErr(msg):
		err = database.ErrFirstNameTaken
	}

	return
}

func (u *Users) GetById(id int) (models.User, error) {
	var user models.User

	row := u.db.QueryRow(SELECT_USER_BY_ID, id)

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, database.ErrNoRows
		}

		return user, internalDatabaseErr(err)
	}

	return user, nil
}

func (u *Users) GetAll() ([]models.User, error) {
	var users []models.User

	rows, err := u.db.Query(SELECT_ALL_USERS)
	if err != nil {
		return nil, internalDatabaseErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.CreatedAt,
			&user.IsActive,
		)
		if err != nil {
			return nil, internalDatabaseErr(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *Users) Create(user models.User) (models.User, error) {
	result, err := u.db.Exec(
		INSERT_USER,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.RoleId,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err = u.getUserErr(sqliteErr.Error())
			return models.User{}, err
		}

		return models.User{}, internalDatabaseErr(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, internalDatabaseErr(err)
	}

	user.Id = int(id)
	return user, nil
}

func (u *Users) Update(user models.User) (models.User, error) {
	_, err := u.db.Exec(
		UPDATE_USER_BY_ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.RoleId,
		user.IsActive,
		user.Id,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err = u.getUserErr(sqliteErr.Error())
			return models.User{}, err
		}

		return models.User{}, internalDatabaseErr(err)
	}

	return user, nil
}
