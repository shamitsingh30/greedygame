package models

import (
	"sync"
)

type Datastore struct {
	Data map[string]string
	sync.RWMutex
}

func NewDatastore() *Datastore {
	return &Datastore{
		Data: map[string]string{},
	}
}
