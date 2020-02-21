package sequego

type modelOptions struct {
	allowNull     bool
	autoIncrement bool
	columnType    string
	primaryKey    bool
}

//Model defines a SQL table and it fields
type Model struct {
	fields map[string]modelOptions
}
