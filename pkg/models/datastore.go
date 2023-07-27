package models

import (
	"github.com/shamitsingh30/greedygame/pkg/types"
)

func NewDatastore() *types.Datastore {
	return &types.Datastore{
		Data: make(map[string]string),
	}
}
