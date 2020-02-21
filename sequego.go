package sequego

import "database/sql"

// ExtendedDB changes the default SQL Type
type extendedDB struct {
	connection *sql.DB
	models     []Model
}
