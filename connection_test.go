package sequego

import (
	"fmt"
	"testing"
)

type databaseConnectionTest struct {
	user       string
	password   string
	dataSource string
}

func TestConnect(t *testing.T) {
	tests := []databaseConnectionTest{
		{user: "root", password: "root", dataSource: "sequego"},
		{user: "root", password: "rot", dataSource: "sequego"},
		{user: "root", password: "root", dataSource: "seque"},
		{user: "", password: "root", dataSource: "seque"},
		{user: "root", password: "", dataSource: "seque"},
		{user: "root", password: "root", dataSource: ""},
	}
	for _, test := range tests {
		testName := fmt.Sprintf("%s %s %s", test.user, test.password, test.dataSource)
		t.Run(testName, func(t *testing.T) {
			err := Connect(test.user, test.password, test.dataSource)

			if err != nil {
				t.Error(err)
			}
		})
	}
}
