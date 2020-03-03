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
	err    error
}

//CreateTable Create a Database Table
func (table Model) CreateTable() error {
	if Connection == nil || Connection.connection == nil {
		return fmt.Errorf("The connection with the database was not initialized yet. Call sequego.Connect(username, password, dataSource) to create a Connection")
	}

	if table.name == "" {
		return fmt.Errorf("You haven't set a name for the model")
	}

	if table.fields == nil {
		return fmt.Errorf("You haven't set any fields for the model")
	}

	primaryKey := 0

	fields := ""
	primaryKeyDefinition := ""

	for field, definition := range table.fields {
		parsedField := parseFields(field, definition)

		if parsedField.primaryKey != "" {
			primaryKey++
			primaryKeyDefinition = parsedField.primaryKey
		}

		fields += parsedField.field
	}

	if primaryKey > 1 {
		return fmt.Errorf("It was not possible to create table due to more than one primary key defined")
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

//AddModelToConnection is a method of the struct Model and adds a model to the existing connection
func (table Model) AddModelToConnection() error {
	if Connection == nil {
		return fmt.Errorf("The connection with the database was not initialized yet. Call sequego.Connect(username, password, dataSource) to create a Connection")
	}

	if table.name == "" {
		return fmt.Errorf("You haven't set a name for the model")
	}

	if table.fields == nil {
		return fmt.Errorf("You haven't set any fields for the model")
	}

	primaryKeyCounter := 0

	for _, definition := range table.fields {
		if definition.primaryKey {
			primaryKeyCounter++
		}
	}

	if primaryKeyCounter > 1 {
		return fmt.Errorf("It was not possible to create table due to more than one primary key defined")
	}

	Connection.models = append(Connection.models, table)

	return nil
}

//AddManyModelsToConnection just calls AddModelToConnection receiving an array
func AddManyModelsToConnection(models []Model) error {
	for _, model := range models {
		err := model.AddModelToConnection()
		if err != nil {
			return fmt.Errorf("It was not possible to associate the model %s to the connection due to %v", model.name, err)
		}
	}
	return nil
}
