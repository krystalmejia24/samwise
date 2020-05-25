package db

import (
	"time"

	"github.com/go-redis/redis"
)

//Svc is the interface declaration
type Svc interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Delete(string) (int64, error)
}

// Redis persists redis client
type Redis struct {
	client *redis.Client
}

// NewRedis initializes and returns the redis client
func NewRedis(addr string) *Redis {
	c := redis.NewClient(&redis.Options{
		Addr: addr, //"redis-server:6379",
	})

	return &Redis{client: c}
}

// Get retrieves value from redis
func (r *Redis) Get(key string) (interface{}, error) {
	val, err := r.client.Get(key).Result()

	return val, err
}

// Set stores value in redis
func (r *Redis) Set(key string, value interface{}) error {
	err := r.client.Set(key, value, 10*time.Minute).Err()

	return err
}

// Delete removes value in redis
func (r *Redis) Delete(key string) (int64, error) {
	val, err := r.client.Del(key).Result()

	return val, err
}
