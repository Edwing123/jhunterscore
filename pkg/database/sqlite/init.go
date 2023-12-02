package sqlite

import _ "embed"

//go:embed base.db
var INITIALIZED_DATABASE_DATA []byte
