package sequego

import (
	"fmt"

	//Getting the mysql definitions
	_ "github.com/go-sql-driver/mysql"
)

//Model defines a SQL table and it fields
type Model struct {
	fields map[string]ModelOptions
	name   string
}

//CreateTable Create a Database Table
//TODO Create tests for this function
func (table Model) CreateTable() error {
	if Connection == nil || Connection.connection == nil {
		return fmt.Errorf(`
			The connection with the database was not initialized yet. 
			Call sequego.Connect(username, password, dataSource) to create a Connection`)
	}

	primaryKey := 0

	fields := ""
	primaryKeyDefinition := ""

	for field, definition := range table.fields {
		if primaryKey > 1 {
			return fmt.Errorf("It was not possible to create table due to more than one primary key defined")
		}
		parsedField := parseFields(field, definition)

		if parsedField.primaryKey {
			primaryKey++
			primaryKeyDefinition = parsedField.primaryKey
		}

		fields += parsedField.field
	}

	fields += primaryKeyDefinition

	statement, err := Connection.connection.Prepare(fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				%s
			);
		`, table.name, fields))

	if err != nil {
		return fmt.Errorf("It was not possible to create table due to %v", err)
	}

	statement.Exec()

	return nil
}
