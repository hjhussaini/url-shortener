package cache

type Cache interface {
	Push(key string, value string) error
	Pop(key string) (string, error)
}
