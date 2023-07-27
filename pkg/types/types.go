package types

import (
	"sync"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type Datastore struct {
	Data map[string]string
	mu   sync.RWMutex
}
