package main

const (
	SESSION_KEY      = "session"
	USER_ROLE_KEY    = "user_role"
	USER_ID_KEY      = "user_id"
	IS_LOGGED_IN_KEY = "is_logged"
	ERRORS_KEY       = "errors"

	DATA_DIR_LOGS_DIR    = "logs"
	DATA_DIR_SESSION_DIR = "sessions"
	DATA_DIR_FILES_DIR   = "files"
	LOGS_FILE            = "logs.txt"
	SESSION_FILE         = "sessions.db"
	DATABASE_FILE        = "jobshunters.db"

	// The file can be of 5MB maximum
	MAX_FILE_SIZE = 1024 * 1024 * 5
)
