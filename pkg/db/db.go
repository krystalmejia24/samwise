package db

import (
	"time"

	"github.com/go-redis/redis"
)

// Repo persists redis client
type Repo struct {
	client *redis.Client
}

// NewRepo initializes and returns the redis client
func NewRepo() *Repo {
	c := redis.NewClient(&redis.Options{
		Addr: "redis-server:6379",
	})

	return &Repo{client: c}
}

// Get retrieves value from redis
func (r *Repo) Get(key string) (interface{}, error) {
	val, err := r.client.Get(key).Result()

	return val, err
}

// Set stores value in redis
func (r *Repo) Set(key string, value interface{}) error {
	err := r.client.Set(key, value, 10*time.Minute).Err()

	return err
}

// Delete removes value in redis
func (r *Repo) Delete(key string) (int64, error) {
	val, err := r.client.Del(key).Result()

	return val, err
}
