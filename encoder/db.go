package encoder

import (
	"encoding/json"

	"github.com/krystalmejia24/samwise/db"
)

type repo struct {
	svc db.Svc
}

const encPrefix string = "encoder:"

//NewRepo returns a repo with a redis connection
func NewRepo(r *db.Redis) *repo {
	return &repo{svc: r}
}

// CreateEncoder sets encoder record in redis
func (r *repo) CreateEncoder(e *Encoder) error {
	key := encPrefix + e.IP
	store, err := json.Marshal(e)
	if err != nil {
		return err
	}
	err = r.svc.Set(key, store)

	return err
}

// GetEncoder gets encoder record from redis
func (r *repo) GetEncoder(ip string) (interface{}, error) {
	key := encPrefix + ip
	val, err := r.svc.Get(key)

	return val, err
}

// DeleteEncoder deletes encoder record from redis
func (r *repo) DeleteEncoder(ip string) (int64, error) {
	key := encPrefix + ip
	val, err := r.svc.Delete(key)

	return val, err
}
