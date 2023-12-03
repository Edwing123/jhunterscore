CREATE TABLE "roles" (
	"role_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	-- Either "cms" or "admin".
	"name" VARCHAR(5) UNIQUE NOT NULL CHECK("name" IN ("cms", "admin"))
);

INSERT INTO "roles" ("name") VALUES
	("admin"),
	("cms");

CREATE TABLE "users" (
	"user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"username" VARCHAR(40) UNIQUE NOT NULL,
	"password" VARCHAR(60) NOT NULL,
	"first_name" VARCHAR(50) UNIQUE NOT NULL,
	"last_name" VARCHAR(50) NOT NULL,
	"role_id" INTEGER NOT NULL REFERENCES "roles"("role_id"),
	-- UTC time zone timestamp.
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_active" TINYINT(1) NOT NULL DEFAULT 1
);

-- Default admin user.
INSERT INTO "users" (
	"username",
	"password",
	"first_name",
	"last_name",
	"role_id"
) VALUES (
	"edwing123",
	"$2a$12$zLR9mjLmDU1WRRw2jjU5xOuOCKW2zEW/YbugayOqdOCRhndg.09Ha",
	"Edwin",
	"Garcia",
	1
);

-- Index for username.
CREATE UNIQUE INDEX "idx_users_username" ON "users"("username");

CREATE TABLE "locations" (
	"location_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(15) UNIQUE NOT NULL
);

-- Populate the table with every city in Nicaragua.
INSERT INTO "locations"("name") VALUES
	("Boaco"),
	("Carazo"),
	("Chinandega"),
	("Chontales"),
	("Estelí"),
	("Granada"),
	("Jinotega"),
	("León"),
	("Madriz"),
	("Managua"),
	("Masaya"),
	("Matagalpa"),
	("Nueva Segovia"),
	("RAAN"),
	("RAAS"),
	("Río San Juan"),
	("Rivas");

CREATE TABLE "files" (
	"file_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	-- A description of the file.
	"name" VARCHAR(100) NOT NULL UNIQUE CHECK(LENGTH("name") > 0),
	-- A reasonable limit for a MIME type is 64 characters long.
	-- files like image/jpeg or text/plain,
	-- etc. 64 characters is more than enough.
	--
	-- A more strong option would be
	-- to have a table that stores all the MIME types
	-- that are supported by the application.
	--
	-- Structure of a MIMEType:
	-- https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types#structure_of_a_mime_type
	"mime_type" VARCHAR(64) NOT NULL CHECK(LENGTH("mime_type") > 0),
	-- A UUID is 36 characters long.
	"path" CHAR(36) NOT NULL UNIQUE CHECK(LENGTH("path") = 36),
	"size" INTEGER NOT NULL CHECK("size" >= 0),
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

CREATE TABLE "companies" (
	"company_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(100) UNIQUE NOT NULL CHECK(LENGTH("name") > 0),
	"logo_url" TEXT NOT NULL UNIQUE CHECK(LENGTH("logo_url") > 0)
);

-- Index for company name.
CREATE UNIQUE INDEX "idx_companies_name" ON "companies"("name");


CREATE TABLE "offers" (
	"offer_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL CHECK(LENGTH("title") > 0),
	"role" VARCHAR(20) NOT NULL CHECK(LENGTH("role") > 0),
	"company_id" INTEGER NOT NULL REFERENCES "companies"("company_id"),
	"content" TEXT NOT NULL CHECK(LENGTH("content") > 0),
	-- Either "pasantia" or "trabajo".
	"contract" VARCHAR(20) NOT NULL CHECK("contract" IN ("pasantia", "trabajo")),
	"salary" INTEGER NOT NULL CHECK("salary" >= 0),
	"contact_info" TEXT NOT NULL CHECK(LENGTH("contact_info") > 0),
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

-- Create index for "role" and "title".
CREATE UNIQUE INDEX "idx_offers_title" ON "offers"("title");
CREATE UNIQUE INDEX "idx_offers_role" ON "offers"("role");

CREATE TABLE "resources" (
	"resource_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL CHECK(LENGTH("title") > 0),
	"summary" VARCHAR(150) NOT NULL CHECK(LENGTH("summary") > 0),
	"content" TEXT NOT NULL CHECK(LENGTH("content") > 0),
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

-- Create index for "title".
CREATE UNIQUE INDEX "idx_resources_title" ON "resources"("title");
