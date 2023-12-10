package sqlite

import (
	"database/sql"
	"errors"
	"strings"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/mattn/go-sqlite3"
)

type Companies struct {
	db *sql.DB
}

func (c *Companies) getCompanyErr(msg string) (err error) {
	switch true {
	case strings.Contains(msg, "name"):
		err = database.ErrCompanyNameTaken
	}

	return
}

func (c *Companies) GetById(id int) (models.Company, error) {
	var company models.Company

	row := c.db.QueryRow(SELECT_COMPANY_BY_ID, id)
	err := row.Scan(
		&company.Id,
		&company.Name,
		&company.LogoURL,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return company, database.ErrNoRows
		}

		return company, internalDatabaseErr(err)
	}

	return company, nil
}

func (c *Companies) GetAll() ([]models.Company, error) {
	companies := []models.Company{}

	rows, err := c.db.Query(SELECT_ALL_COMPANIES)
	if err != nil {
		return nil, internalDatabaseErr(err)
	}

	for rows.Next() {
		var company models.Company

		err := rows.Scan(
			&company.Id,
			&company.Name,
			&company.LogoURL,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, database.ErrNoRows
			}

			return nil, internalDatabaseErr(err)
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (c *Companies) Create(company models.Company) (models.Company, error) {
	result, err := c.db.Exec(
		INSERT_COMPANY,
		company.Name,
		company.LogoURL,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err := c.getCompanyErr(sqliteErr.Error())
			return models.Company{}, err
		}

		return models.Company{}, internalDatabaseErr(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Company{}, internalDatabaseErr(err)
	}

	company.Id = int(id)
	return company, nil
}

func (c *Companies) Update(company models.Company) (models.Company, error) {
	_, err := c.db.Exec(
		UPDATE_COMPANY_BY_ID,
		company.Name,
		company.LogoURL,
		company.Id,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err := c.getCompanyErr(sqliteErr.Error())
			return models.Company{}, err
		}

		return models.Company{}, internalDatabaseErr(err)
	}

	return company, nil
}

func (c *Companies) Delete(id int) error {
	_, err := c.db.Exec(DELETE_COMPANY_BY_ID, id)
	if err != nil {
		return internalDatabaseErr(err)
	}

	return nil
}
