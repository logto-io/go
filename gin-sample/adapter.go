package main

import (
	"github.com/gin-contrib/sessions"
)

type Storage struct {
	session sessions.Session
}

func (storage *Storage) GetItem(key string) string {
	value := storage.session.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (storage *Storage) SetItem(key, value string) {
	storage.session.Set(key, value)
	storage.session.Save()
}
