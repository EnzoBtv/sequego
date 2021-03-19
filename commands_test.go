package sequego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type exampleModel struct {
	id   int
	name string
}

type selectAllColumnsTest struct {
	inModel  Model
	extModel exampleModel
	err      error
}

func TestSelectAllColumns(t *testing.T) {
	tests := map[string]selectAllColumnsTest{
		"Must pass test": {
			inModel: Model{name: "Test1",
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
				},
			},
			extModel: exampleModel{},
		},
	}

	// Connect("root", "root", "sequego", "localhost", 3306)

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			err := test.inModel.SelectAllColumns(&test.extModel)
			if test.err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.err == nil && err != nil {
				t.Errorf("Test failed due to %v", err)
			}
		})
	}
}
