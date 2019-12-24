package db

import (
	//"github.com/google/uuid"
	"encoding/json"
)

const encPrefix string = "encoder:"

// CreateEncoder encoder record in repo
func (r *Repo) CreateEncoder(e *Encoder) error {
	key := encPrefix + e.IP
	store, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = r.Set(key, store)

	return err
}

// GetEncoder gets encoder record from repo
func (r *Repo) GetEncoder(ip string) (interface{}, error) {
	key := encPrefix + ip
	val, err := r.Get(key)

	return val, err
}

// DeleteEncoder deletes encoder record from repo
func (r *Repo) DeleteEncoder(ip string) (int64, error) {
	key := encPrefix + ip
	val, err := r.Delete(key)

	return val, err
}
