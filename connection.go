package sequego

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	//Getting the mysql definitions
	_ "github.com/go-sql-driver/mysql"
)

// Connection stores a connection with a SQL Database
var Connection ExtendedDB

// Connect connects with the specified database
func (db extendedDB) Connect(username, password, dataSource string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", username, password, dataSource)
	connection, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("\nIt was not possible to connect with sql due to %v\n", err)
	}

	Connection = ExtendedDB{
		connection: connection,
		models: [],
	}

	log.Printf("Connection to database %s made successfully", strings.Split(connectionString, "/")[1])
}
