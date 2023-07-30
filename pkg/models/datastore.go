package models

import (
	"sync"
	"time"
)

type Token struct {
	Value      string
	Expiration time.Time
}

type Datastore struct {
	Data map[string]Token
	sync.RWMutex
}

func NewDatastore() *Datastore {
	return &Datastore{
		Data: map[string]Token{},
	}
}
