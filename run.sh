# Run the server

BUILD_TAGS="sqlite_fts5 sqlite_foreign_keys"

go run -tags "$BUILD_TAGS" ./cmd/core -config=./config.exampe.json
