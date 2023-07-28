package controllers

import (
	"errors"
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

func Pop_controller(body *map[string]string, qb *(models.Queuestore)) (string, error) {
	key := (*body)["key"]

	qb.Lock()
	defer qb.Unlock()

	if len(qb.Data[key]) > 0 {
		poppedElement := qb.Data[key][len(qb.Data[key])-1]
		qb.Data[key] = qb.Data[key][:len(qb.Data[key])-1]

		if len(qb.Data[key]) == 0 {
			delete(qb.Data, key)
		}
		return poppedElement, nil
	}
	return "", errors.New("queue is empty")
}
