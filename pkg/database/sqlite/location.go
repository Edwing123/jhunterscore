package sqlite

import (
	"database/sql"

	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
)

type Locations struct {
	db *sql.DB
}

func (l *Locations) GetAll() ([]models.Location, error) {
	locations := []models.Location{}

	rows, err := l.db.Query(SELECT_ALL_LOCATIONS)
	if err != nil {
		return nil, internalDatabaseErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		var location models.Location

		err := rows.Scan(
			&location.Id,
			&location.Name,
		)
		if err != nil {
			return nil, internalDatabaseErr(err)
		}

		locations = append(locations, location)
	}

	return locations, nil
}
