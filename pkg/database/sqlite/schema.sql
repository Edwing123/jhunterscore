CREATE DATABASE jobshunters

CREATE TABLE "users" (
	"user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"username" VARCHAR(40) UNIQUE NOT NULL,
	"password" VARCHAR(60) NOT NULL,
	"first_name" VARCHAR(50) UNIQUE NOT NULL,
	"last_name" VARCHAR(50) UNIQUE NOT NULL,
	-- UTC time zone timestamp.
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_active" TINYINT(1) NOT NULL DEFAULT 1
);

-- Index for username.
CREATE UNIQUE INDEX "idx_users_username" ON "users"("username");

CREATE TABLE "locations" (
	"location_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"name" VARCHAR(50) UNIQUE NOT NULL,
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


CREATE TABLE "offers" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL,
	-- Either "admin" or "cms".
	"role" VARCHAR(20) NOT NULL CHECK("role" IN ("admin", "cms")),
	"company" VARCHAR(100) NOT NULL,
	"content" TEXT NOT NULL,
	-- Either "pasantia" or "trabajo".
	"contract" VARCHAR(20) NOT NULL CHECK("contract" IN ("pasantia", "trabajo")),
	"salary" INTEGER NOT NULL CHECK("salary" >= 0),
	"contact_info" TEXT NOT NULL,
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);

CREATE TABLE "resources" (
	"resource_id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL,
	"summary" VARCHAR(150) NOT NULL,
	"content" TEXT NOT NULL,
	"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"is_published" TINYINT(1) NOT NULL DEFAULT 0,
	"user_id" INTEGER NOT NULL REFERENCES "users"("user_id")
);
