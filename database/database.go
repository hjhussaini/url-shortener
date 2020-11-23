package database

type Session interface {
	Count(table string) (int64, error)
	Insert(
		ttl int64,
		table string,
		fields string,
		values ...interface{},
	) error
}
