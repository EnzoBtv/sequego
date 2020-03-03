package sequego

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {
	tests := map[string]Model{
		"Must Pass Test": {
			name: "Test1",
			fields: map[string]ModelOptions{
				"id": {
					allowNull:     false,
					autoIncrement: true,
					primaryKey:    true,
					columnType:    "INT",
				},
				"name": {
					allowNull:     false,
					autoIncrement: false,
					primaryKey:    false,
					columnType:    "VARCHAR(255)",
				},
				"address": {
					allowNull:     true,
					autoIncrement: false,
					primaryKey:    false,
					columnType:    "VARCHAR(255)",
				},
			},
		},
		"Fields Missing Test": {
			name: "Test1",
			err:  fmt.Errorf("You haven't set any fields for the model"),
		},
		"Duplicate Primary Key": {
			name: "Test1",
			fields: map[string]ModelOptions{
				"id": {
					allowNull:     false,
					autoIncrement: true,
					primaryKey:    true,
					columnType:    "INT",
				},
				"name": {
					allowNull:     false,
					autoIncrement: false,
					primaryKey:    true,
					columnType:    "VARCHAR(255)",
				},
				"address": {
					allowNull:     true,
					autoIncrement: false,
					primaryKey:    false,
					columnType:    "VARCHAR(255)",
				},
			},
			err: fmt.Errorf("It was not possible to create table due to more than one primary key defined"),
		},
	}

	Connect("root", "root", "sequego", "localhost", 3306)

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			err := test.CreateTable()
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.err == nil && err != nil {
				t.Errorf("Test failed due to %v", err)
			}
		})
	}
}
