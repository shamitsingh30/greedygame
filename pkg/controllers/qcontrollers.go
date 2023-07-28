package controllers

import (
	"strings"

	"github.com/shamitsingh30/greedygame/pkg/models"
)

func Push_controller(body *map[string]string, qb *(models.Queuestore)) {

	key := (*body)["key"]

	qb.Lock()
	defer qb.Unlock()
	_, exist := qb.Data[key]

	if !exist {
		qb.Data[key] = make([]string, 0)
	}
	items := strings.Split((*body)["items"], " ")

	qb.Data[key] = append(qb.Data[key], items...)

	return
}
