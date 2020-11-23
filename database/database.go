package database

type Session interface {
	Count(table string) (int64, error)
	Select(
		table string,
		fields string,
		condition string,
		limit int,
	) SelectResult
	Insert(
		ttl int64,
		table string,
		fields string,
		values ...interface{},
	) error
	Delete(table string, fields_values map[string]string) error
}

type SelectResult interface {
	Scan(values ...interface{}) bool
}
