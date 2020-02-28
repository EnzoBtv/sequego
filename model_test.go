package sequego

import (
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
	}

	Connect("root", "root", "sequego")

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
