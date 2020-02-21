package sequego

// ExtendedDB changes the default SQL Type
type ExtendedDB struct {
	connection *sql.DB
	models     []Model
}
/