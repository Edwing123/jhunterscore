package sqlite

import (
	"database/sql"
	"errors"
	"strings"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"github.com/mattn/go-sqlite3"
)

type Offers struct {
	db *sql.DB
}

func (f *Offers) isTitleTakenErr(msg string) bool {
	return strings.Contains(msg, "title")
}

func (f *Offers) getOfferErr(msg string) (err error) {
	switch true {
	case f.isTitleTakenErr(msg):
		err = database.ErrOfferTitleTaken
	}

	return
}

func (o *Offers) GetById(id int) (models.Offer, error) {
	var offer models.Offer

	row := o.db.QueryRow(SELECT_OFFER_BY_ID, id)

	err := row.Scan(
		&offer.Id,
		&offer.Title,
		&offer.Role,
		&offer.Company,
		&offer.CompanyLogoURL,
		&offer.Content,
		&offer.Contract,
		&offer.Location,
		&offer.Salary,
		&offer.ContactInfo,
		&offer.CreatedAt,
		&offer.IsPublished,
		&offer.Author,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return offer, database.ErrNoRows
		}

		return offer, internalDatabaseErr(err)
	}

	return offer, nil
}

func (o *Offers) GetUserIdByOfferId(id int) (int, error) {
	var userId int

	row := o.db.QueryRow(SELECT_OFFER_USER_ID_BY_ID, id)

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

func (o *Offers) GetAll() ([]models.Offer, error) {
	offers := []models.Offer{}

	rows, err := o.db.Query(SELECT_ALL_OFFERS)
	if err != nil {
		return nil, internalDatabaseErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		var offer models.Offer

		err := rows.Scan(
			&offer.Id,
			&offer.Title,
			&offer.Role,
			&offer.Company,
			&offer.CompanyLogoURL,
			&offer.Content,
			&offer.Contract,
			&offer.Location,
			&offer.Salary,
			&offer.ContactInfo,
			&offer.CreatedAt,
			&offer.IsPublished,
			&offer.Author,
		)
		if err != nil {
			return nil, internalDatabaseErr(err)
		}

		offers = append(offers, offer)
	}

	return offers, nil
}

func (o *Offers) Create(offer models.Offer) (models.Offer, error) {
	result, err := o.db.Exec(
		INSERT_OFFER,
		offer.Title,
		offer.Role,
		offer.CompanyId,
		offer.Content,
		offer.Contract,
		offer.LocationId,
		offer.Salary,
		offer.ContactInfo,
		offer.UserId,
	)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && isUniqueViolationErr(sqliteErr) {
			err = o.getOfferErr(sqliteErr.Error())
			return models.Offer{}, err
		}

		return models.Offer{}, internalDatabaseErr(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Offer{}, internalDatabaseErr(err)
	}

	offer.Id = int(id)
	return offer, nil
}

func (o *Offers) Update(offer models.Offer) (models.Offer, error) {
	return offer, nil
}

func (o *Offers) Delete(id int) error {
	_, err := o.db.Exec(DELETE_OFFER_BY_ID, id)
	if err != nil {
		return internalDatabaseErr(err)
	}

	return nil
}
