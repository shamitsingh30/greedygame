package controllers

import (
	"time"

	"github.com/shamitsingh30/greedygame/pkg/models"
)

func Set_controller(body *map[string]string, db *(models.Datastore)) {

	db.Lock()
	key := (*body)["key"]
	value := (*body)["value"]

	condition, exist := (*body)["condition"]
	if exist {
		_, ok := db.Data[key]
		if (condition == "NX" && ok) || (condition == "XX" && !ok) {
			db.Unlock()
			return
		}
	}
	db.Data[key] = value
	db.Unlock()

	expiry, exist := (*body)["expiry_time"]

	if exist {
		expiry, _ := time.ParseDuration(expiry)

		time.AfterFunc(expiry, func() {
			db.Lock()
			delete(db.Data, key)
			db.Unlock()
		})
	}
}
