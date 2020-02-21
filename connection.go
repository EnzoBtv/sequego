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
var Connection extendedDB

// Connect connects with the specified database
func Connect(username, password, dataSource string) error {
	connectionString := fmt.Sprintf("%s:%s@/%s", username, password, dataSource)
	connection, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatalf("\nIt was not possible to connect with sql due to %v\n", err)
		return err
	}

	Connection = extendedDB{
		connection: connection,
	}

	log.Printf("Connection to database %s made successfully", strings.Split(connectionString, "/")[1])

	return nil
}
