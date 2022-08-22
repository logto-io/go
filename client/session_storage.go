package client

type SessionStorage interface {
	GetItem(key string) string
	SetItem(key, value string)
}
