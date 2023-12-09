package sqlite

const (
	// Auth and User queries.
	SELECT_USER_BY_FOR_AUTH = `SELECT "user_id", "password" FROM "users" WHERE "username" = ?;`

	SELECT_ALL_USERS = `
	SELECT "user_id", "username", "first_name", "last_name", "r"."name" AS "role", "created_at", "is_active"
	FROM "users" AS "u"
	INNER JOIN "roles" AS "r"
	ON "u"."role_id" = "r"."role_id";
	`

	SELECT_USER_BY_ID = `
	SELECT "user_id", "username", "first_name", "last_name", "r"."name" AS "role", "created_at", "is_active"
	FROM "users" AS "u'
	INNER JOIN "roles" AS "r"
	ON "u"."role_id" = "r"."role_id"
	WHERE "id" = ?;
	`

	INSERT_USER = `
	INSERT INTO "users" ("username", "password", "first_name", "last_name", "role_id")
	VALUES (?, ?, ?, ?, ?);
	`

	UPDATE_USER_BY_ID = `
	UPDATE "users"
	SET	"username" = ?,
		"first_name" = ?,
		"last_name" = ?,
		"role_id" = ?,
		"is_active" = ?
	WHERE "user_id" = ?;
	`

	// File queries.
	SELECT_FILE_BY_ID = `
	SELECT "file_id", "name", "mime_type", "path", "size", ("u"."first_name" || ' ' || "u"."last_name") AS "author", "created_at"
	FROM "files" AS "f"
	INNER JOIN "users" AS "u"
	ON "f"."user_id" = "u"."user_id"
	WHERE "file_id" = ?;
	`

	SELECT_FILE_USER_ID_BY_ID = `
	SELECT "user_id"
	FROM "files"
	WHERE "file_id" = ?;
	`

	SELECT_ALL_FILES = `
	SELECT "file_id", "name", "mime_type", "path", "size", ("u"."first_name" || ' ' || "u"."last_name") AS "author", "created_at"
	FROM "files" AS "f"
	INNER JOIN "users" AS "u"
	ON "f"."user_id" = "u"."user_id";
	`

	INSERT_FILE = `
	INSERT INTO "files" ("name", "mime_type", "path", "size", "user_id")
	VALUES (?, ?, ?, ?, ?);
	`

	DELETE_FILE_BY_ID = `
	DELETE FROM "files"
	WHERE "file_id" = ?;
	`

	UPDATE_FILE_BY_ID = `
	UPDATE "files"
	SET "name" = ?
	WHERE "file_id" = ?;
	`

	// Company queries.
	SELECT_COMPANY_BY_ID = `
	SELECT "company_id", "name", "logo_url"
	FROM "companies"
	WHERE "company_id" = ?;
	`

	SELECT_ALL_COMPANIES = `
	SELECT "company_id", "name", "logo_url"
	FROM "companies";
	`

	INSERT_COMPANY = `
	INSERT INTO "companies" ("name", "logo_url")
	VALUES(?, ?);
	`

	UPDATE_COMPANY_BY_ID = `
	UPDATE "companies"
	SET "name" = ?,
		"logo_url" = ?
	WHERE "company_id" = ?;
	`

	DELETE_COMPANY_BY_ID = `
	DELETE FROM "companies"
	WHERE "company_id" = ?;
	`
)
