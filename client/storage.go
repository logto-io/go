package client

type Storage interface {
	GetItem(key string) string
	SetItem(key, value string)
}
