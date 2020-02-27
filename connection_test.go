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
	err        error
}

func TestConnect(t *testing.T) {
	tests := []databaseConnectionTest{
		{user: "root", password: "root", dataSource: "sequego"},
		{user: "root", password: "rot", dataSource: "sequego"},
		{user: "root", password: "root", dataSource: "seque"},
		{user: "", password: "root", dataSource: "seque", err: fmt.Errorf("The username (first parameter) is required")},
		{user: "root", password: "", dataSource: "seque"},
		{user: "root", password: "root", dataSource: "", err: fmt.Errorf("The datasource (third parameter) is required")},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("%s %s %s", test.user, test.password, test.dataSource)
		t.Run(testName, func(t *testing.T) {
			err := Connect(test.user, test.password, test.dataSource)

			if test.err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.err == nil && err != nil {
				t.Errorf("Test failed due to %v", err)
			}
		})
	}
}
