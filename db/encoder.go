package db

import (
	"encoding/json"
)

const encPrefix string = "encoder:"

// CreateEncoder sets encoder record in redis
func CreateEncoder(e *Encoder, repo Svc) error {
	key := encPrefix + e.IP
	store, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = repo.Set(key, store)

	return err
}

// GetEncoder gets encoder record from redis
func GetEncoder(ip string, repo Svc) (interface{}, error) {
	key := encPrefix + ip
	val, err := repo.Get(key)

	return val, err
}

// DeleteEncoder deletes encoder record from redis
func DeleteEncoder(ip string, repo Svc) (int64, error) {
	key := encPrefix + ip
	val, err := repo.Delete(key)

	return val, err
}
