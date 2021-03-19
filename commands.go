package sequego

import (
	"fmt"
	"reflect"
	"strings"
)

func (table *Model) SelectAllColumns(dataReceiver interface{}) error {
	if dataReceiver == nil {
		return fmt.Errorf("Datareceiver parameter cannot be null")
	}

	selectQuery := `SELECT`

	for column := range table.fields {
		selectQuery += fmt.Sprintf(" %s,", column)
	}
	selectQuery = strings.TrimSuffix(selectQuery, ",")

	selectQuery += fmt.Sprintf(" FROM %s", table.name)
	rows, err := connection.connection.Query(selectQuery)
	if err != nil {
		return fmt.Errorf("Unexpected error on SELECT, please check your model fields")
	}

	for rows.Next() {
		var test interface{}
		err := rows.Scan(test)
		fmt.Println(err)
	}

	dataReceiverReflected := reflect.ValueOf(dataReceiver)
	dataReceiverReflectedPointer := reflect.ValueOf(&dataReceiver)
	dataReceiverType := dataReceiverReflected.Type()
	dataReceiverElement := dataReceiverReflectedPointer.Elem()

	tmp := reflect.New(dataReceiverElement.Elem().Type()).Elem()
	tmp.Set(dataReceiverElement.Elem())

	for i := 0; i < dataReceiverReflected.NumField(); i++ {
		field := dataReceiverType.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "-" {
			tmp.FieldByName(field.Name)
		}
	}

	return nil
}
