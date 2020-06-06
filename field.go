package sequego

import (
	"fmt"
)

//ModelOptions defines how a SQL field can be
type ModelOptions struct {
	allowNull     bool
	autoIncrement bool
	columnType    string
	primaryKey    bool
}

type Field struct {
	primaryKey string
	field      string
}

func (p *Field) parse(field string, definition ModelOptions) {
	fieldDefinition := fmt.Sprintf("%s %s", field, definition.columnType)

	if !definition.allowNull {
		fieldDefinition += " NOT NULL"
	}

	if definition.autoIncrement {
		fieldDefinition += " AUTO_INCREMENT"
	}

	if definition.primaryKey {
		p.primaryKey = fmt.Sprintf("PRIMARY KEY (%s)", field)
	}
	fieldDefinition += ",\n"
	p.field = fieldDefinition
}
