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

type returnParseFields struct {
	primaryKey string
	field      string
}

func parseFields(field string, definition ModelOptions) returnParseFields {
	returnable := &returnParseFields{}

	fieldDefinition := fmt.Sprintf("%s %s", field, definition.columnType)

	if !definition.allowNull {
		fieldDefinition += " NOT NULL"
	}

	if definition.autoIncrement {
		fieldDefinition += " AUTO_INCREMENT"
	}

	if definition.primaryKey {
		returnable.primaryKey = fmt.Sprintf("PRIMARY KEY (%s)", field)
	}
	fieldDefinition += ",\n"
	returnable.field = fieldDefinition

	return *returnable
}
