package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/shamitsingh30/greedygame/pkg/models"
)

func Set_controller(body *map[string]string, db *(models.Datastore)) {

	key := (*body)["key"]
	val := (*body)["value"]

	condition, exist := (*body)["condition"]
	if exist {
		_, _, err := Get_controller(body, db)
		if (err == nil && condition == "NX") || (err != nil && condition == "XX") {
			return
		}
	}

	e := models.Token{
		Value: val,
	}
	expiry, exist := (*body)["expiry_time"]

	if exist {
		expiryTime, _ := strconv.Atoi(expiry)
		e.Expiration = time.Now().Add(time.Duration(expiryTime) * time.Second)
	}
	db.Lock()
	defer db.Unlock()
	db.Data[key] = e
	return
}

func Get_controller(body *map[string]string, db *(models.Datastore)) (string, string, error) {

	db.Lock()
	defer db.Unlock()

	key := (*body)["key"]
	val, exists := db.Data[key]

	if exists {
		if !val.Expiration.IsZero() {
			if db.Data[key].Expiration.Before(time.Now()) {
				delete(db.Data, key)
			} else {
				return key, val.Value, nil
			}
		} else {
			return key, val.Value, nil
		}
	}

	return "", "", errors.New("key not found")
}
