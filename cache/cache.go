package cache

type Cache interface {
	Push(value ...interface{}) error
	Pop() (string, error)
}
