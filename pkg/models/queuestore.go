package models

import (
	"sync"
)

type Queue struct {
	Items string
}

type Queuestore struct {
	Data map[string][]string
	sync.RWMutex
}

func NewQueuestore() *Queuestore {
	return &Queuestore{
		Data: map[string][]string{},
	}
}
