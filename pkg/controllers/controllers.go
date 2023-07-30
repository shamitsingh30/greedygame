package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/shamitsingh30/greedygame/pkg/models"
)

func check_existence(key string, db *(models.Datastore)) bool {
	val, exists := db.Data[key]

	if exists && (val.Expiration.IsZero() || val.Expiration.After(time.Now())) {
		return true
	}
	return false
}

func Set_controller(body *map[string]string, db *(models.Datastore)) {

	key := (*body)["key"]
	val := (*body)["value"]
	condition, exist := (*body)["condition"]

	db.Lock()
	defer db.Unlock()

	if exist {
		ch := check_existence(key, db)

		if (ch == true && condition == "NX") || (ch == false && condition == "XX") {
			return
		}
	}

	e := models.Token{
		Value: val,
	}
	expiry, ex := (*body)["expiry_time"]

	if ex {
		expiryTime, _ := strconv.Atoi(expiry)
		e.Expiration = time.Now().Add(time.Duration(expiryTime) * time.Second)
	}

	db.Data[key] = e
	return
}

func Get_controller(body *map[string]string, db *(models.Datastore)) (string, string, error) {

	key := (*body)["key"]

	db.RLock()
	val, exists := db.Data[key]
	db.RUnlock()

	if !exists {
		return "", "", errors.New("key not found")
	}

	if !val.Expiration.IsZero() && val.Expiration.Before(time.Now()) {
		db.Lock()
		delete(db.Data, key)
		db.Unlock()
		return "", "", errors.New("key not found")
	}

	return key, val.Value, nil
}
