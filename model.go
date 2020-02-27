package sequego

import (
	"fmt"

	//Getting the mysql definitions
	_ "github.com/go-sql-driver/mysql"
)

//ModelOptions defines how a SQL field can be
type ModelOptions struct {
	allowNull     bool
	autoIncrement bool
	columnType    string
	primaryKey    bool
}

//Model defines a SQL table and it fields
type Model struct {
	fields map[string]ModelOptions
	name   string
}

//CreateTable Create a Database Table
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

		fieldDefinition := fmt.Sprintf("%s %s", field, definition.columnType)

		if !definition.allowNull {
			fieldDefinition += " NOT NULL"
		}

		if definition.autoIncrement {
			fieldDefinition += " AUTO_INCREMENT"
		}

		if definition.primaryKey {
			primaryKey++
			primaryKeyDefinition = fmt.Sprintf("PRIMARY KEY (%s)", field)
		}
		fieldDefinition += "\n"
		fields += fieldDefinition
	}

	fields += primaryKeyDefinition

	Connection.connection.Prepare(fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				%s
			);
		`, table.name, fields))

	return nil
}
