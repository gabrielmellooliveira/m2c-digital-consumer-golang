package database

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisAdapter struct {
	redisClient *redis.Client
}

func NewRedisAdapter(url string) *RedisAdapter {
	client := redis.NewClient(&redis.Options{
		Addr: url,
	})

	return &RedisAdapter{
		redisClient: client,
	}
}

func (r *RedisAdapter) Connect() error {
	_, err := r.redisClient.Ping(context.Background()).Result()
	if err != nil {
		return errors.New("Failed to connect to Redis: " + err.Error())
	}
	return nil
}

func (r *RedisAdapter) Disconnect() error {
	err := r.redisClient.Close()
	if err != nil {
		return errors.New("Failed to disconnect to Redis: " + err.Error())
	}
	return nil
}

func (r *RedisAdapter) Set(key string, value interface{}) error {
	err := r.redisClient.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return errors.New("Failed to set value in Redis: " + err.Error())
	}
	return nil
}

func (r *RedisAdapter) Get(key string) (string, error) {
	value, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.New("Failed to get value from Redis: " + err.Error())
	}
	return value, nil
}

func (r *RedisAdapter) Increment(key string) error {
	err := r.redisClient.Incr(context.Background(), key).Err()
	if err != nil {
		return errors.New("Failed to increment key in Redis: " + err.Error())
	}
	return nil
}

func (r *RedisAdapter) Delete(key string) error {
	err := r.redisClient.Del(context.Background(), key).Err()
	if err != nil {
		return errors.New("Failed to delete key in Redis: " + err.Error())
	}
	return nil
}
