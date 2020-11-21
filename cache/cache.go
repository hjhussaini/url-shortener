package cache

type Cache interface {
	Push(value string) error
	Pop() (string, error)
}
