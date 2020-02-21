package sequego

import (
	"database/sql"
	"fmt"
	"log"

	//Getting the mysql definitions
	_ "github.com/go-sql-driver/mysql"
)

type modelOptions struct {
	allowNull     bool
	autoIncrement bool
	columnType    string
	primaryKey    bool
}

type model struct {
	fields map[string]modelOptions
}

type extendedDB struct {
	connection *sql.DB
	models     []model
}

// Connection stores a connection with a SQL Database
var Connection *sql.DB

// Connect connects with the specified database
func Connect(username, password, dataSource string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", username, password, dataSource)
	connection, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("\nIt was not possible to connect with sql due to %v\n", err)
	}

	Connection = connection
}
