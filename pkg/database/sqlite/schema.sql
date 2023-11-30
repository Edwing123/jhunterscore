CREATE DATABASE "jobshunters"

CREATE TABLE "roles" (
	"role_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	-- Either "cms" or "admin".
	"name" VARCHAR(5) UNIQUE NOT NULL CHECK("name" IN ("cms", "admin"))
);

INSERT INTO "roles" ("name") VALUES
	("cms"),
	("admin");

CREATE TABLE "users" (
	"user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"username" VARCHAR(40) UNIQUE NOT NULL,
	"password" VARCHAR(60) NOT NULL,
	"first_name" VARCHAR(50) UNIQUE NOT NULL,
	"last_name" VARCHAR(50) UNIQUE NOT NULL,
	"role_id" INTEGER NOT NULL REFERENCES "roles"("role_id"),
	-- UTC time zone timestamp.
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_active" TINYINT(1) NOT NULL DEFAULT 1
);

-- Index for username.
CREATE UNIQUE INDEX "idx_users_username" ON "users"("username");

CREATE TABLE "locations" (
	"location_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(15) UNIQUE NOT NULL,
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

CREATE TABLE "companies" (
	"company_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(100) UNIQUE NOT NULL
);

-- Index for company name.
CREATE UNIQUE INDEX "idx_companies_name" ON "companies"("name");


CREATE TABLE "offers" (
	"offer_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL,
	"role" VARCHAR(20) NOT NULL,
	"company_id" INTEGER NOT NULL REFERENCES "companies"("company_id"),
	"content" TEXT NOT NULL,
	-- Either "pasantia" or "trabajo".
	"contract" VARCHAR(20) NOT NULL CHECK("contract" IN ("pasantia", "trabajo")),
	"salary" INTEGER NOT NULL CHECK("salary" >= 0),
	"contact_info" TEXT NOT NULL,
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

-- Create index for "role" and "title".
CREATE UNIQUE INDEX "idx_offers_title" ON "offers"("title");
CREATE UNIQUE INDEX "idx_offers_role" ON "offers"("role");

CREATE TABLE "resources" (
	"resource_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL,
	"summary" VARCHAR(150) NOT NULL,
	"content" TEXT NOT NULL,
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

-- Create index for "title".
CREATE UNIQUE INDEX "idx_resources_title" ON "resources"("title");
