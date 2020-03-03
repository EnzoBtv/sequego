package sequego

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type databaseConnectionTest struct {
	user       string
	password   string
	dataSource string
	host       string
	port       int
	err        error
}

func TestConnect(t *testing.T) {
	tests := []databaseConnectionTest{
		{user: "root", password: "root", dataSource: "sequego", host: "127.0.0.1", port: 3306},
		{user: "root", password: "rot", dataSource: "sequego", host: "127.0.0.1", port: 3306, err: fmt.Errorf("The access for the database has been denied, check your connection")},
		{user: "root", password: "root", dataSource: "seque", host: "127.0.0.1", port: 3306, err: fmt.Errorf("The access for the database has been denied, check your connection")},
		{user: "", password: "root", dataSource: "seque", host: "127.0.0.1", port: 3306, err: fmt.Errorf("The username (first parameter) is required")},
		{user: "root", password: "", dataSource: "seque", host: "127.0.0.1", port: 3306, err: fmt.Errorf("The access for the database has been denied, check your connection")},
		{user: "root", password: "root", dataSource: "", host: "127.0.0.1", port: 3306, err: fmt.Errorf("The datasource (third parameter) is required")},
		{user: "root", password: "root", dataSource: "sequego", host: "", port: 3306},
		{user: "root", password: "root", dataSource: "sequego", host: "", port: 0},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("%s %s %s", test.user, test.password, test.dataSource)
		t.Run(testName, func(t *testing.T) {
			err := Connect(test.user, test.password, test.dataSource, test.host, test.port)

			if test.err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.err == nil && err != nil {
				t.Errorf("Test failed due to %v", err)
			}
		})
	}
}
