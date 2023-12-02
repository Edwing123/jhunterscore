package sqlite

import (
	"database/sql"
	"errors"
	"strings"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/mattn/go-sqlite3"
)

type Files struct {
	db *sql.DB
}

func (f *Files) isNameTakenErr(msg string) bool {
	return strings.Contains(msg, "name")
}

func (f *Files) isPathTakenErr(msg string) bool {
	return strings.Contains(msg, "path")
}

func (f *Files) getFileErr(msg string) (err error) {
	switch true {
	case f.isNameTakenErr(msg):
		err = database.ErrFileNameTaken
	case f.isPathTakenErr(msg):
		err = database.ErrFilePathTaken
	}

	return
}

func (f *Files) GetById(id int) (models.File, error) {
	var file models.File

	row := f.db.QueryRow(SELECT_FILE_BY_ID, id)

	err := row.Scan(
		&file.Id,
		&file.Name,
		&file.MIMEType,
		&file.Path,
		&file.Author,
		&file.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return file, database.ErrNoRows
		}

		return file, internalDatabaseErr(err)
	}

	return file, nil
}

func (f *Files) GetUserIdById(id int) (int, error) {
	var userId int

	row := f.db.QueryRow(SELECT_FILE_USER_ID_BY_ID, id)

	err := row.Scan(
		&userId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, database.ErrNoRows
		}

		return 0, internalDatabaseErr(err)
	}

	return userId, nil
}

func (f *Files) GetAll() ([]models.File, error) {
	var files []models.File

	rows, err := f.db.Query(SELECT_ALL_FILES)
	if err != nil {
		return nil, internalDatabaseErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		var file models.File

		err := rows.Scan(
			&file.Id,
			&file.Name,
			&file.MIMEType,
			&file.Path,
			&file.Author,
			&file.CreatedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, database.ErrNoRows
			}

			return nil, internalDatabaseErr(err)
		}

		files = append(files, file)
	}

	return files, nil
}

func (f *Files) Create(file models.File) (models.File, error) {
	result, err := f.db.Exec(
		INSERT_FILE,
		file.Name,
		file.MIMEType,
		file.Path,
		file.UserId,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err = f.getFileErr(sqliteErr.Error())
			return models.File{}, err
		}

		return models.File{}, internalDatabaseErr(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.File{}, internalDatabaseErr(err)
	}

	file.Id = int(id)
	return file, nil
}

func (f *Files) Update(file models.File) (models.File, error) {
	_, err := f.db.Exec(DELETE_FILE_BY_ID, file.Name, file.Id)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err = f.getFileErr(sqliteErr.Error())
			return models.File{}, err
		}

		return models.File{}, internalDatabaseErr(err)
	}

	return file, nil
}

func (f *Files) Delete(id int) error {
	_, err := f.db.Exec(DELETE_FILE_BY_ID, id)
	if err != nil {
		return internalDatabaseErr(err)
	}

	return nil
}
