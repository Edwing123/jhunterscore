package database

import "errors"

var (
	// Authentication errors.
	ErrAuth = errors.New("db: auth failure")

	// Result sets errors.
	ErrNoRows = errors.New("db: no rows in result set")

	// Internal database errors.
	ErrInternalDatabase = errors.New("db: interal database failure")

	// User errors.
	ErrUsernameTaken  = errors.New("db: username taken")
	ErrFirstNameTaken = errors.New("db: first name taken")

	// File errors.
	ErrFileNameTaken = errors.New("db: file name taken")
	ErrFilePathTaken = errors.New("db: file path taken")

	// Company errors.
	ErrCompanyNameTaken = errors.New("db: company name taken")

	// Offer errors.
	ErrOfferTitleTaken = errors.New("db: offer title taken")
)
