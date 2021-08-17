package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Retrieve(key string) (*string, error)

	Put(key string, value string) error

	Delete(key string) error
}

type RedisCache struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedisCache() Cache {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisCache{
		ctx: context.Background(),
		rdb: rdb,
	}

}

func (r *RedisCache) Retrieve(key string) (*string, error) {

	result, error := r.rdb.Get(r.ctx, key).Result()

	if error == redis.Nil {

		return nil, nil

	}

	if error != nil {

		return nil, error

	}

	return &result, nil

}

func (r *RedisCache) Put(key string, value string) error {

	err := r.rdb.Set(r.ctx, key, value, 0).Err()

	return err

}

func (r *RedisCache) Delete(key string) error {

	_, err := r.rdb.Del(r.ctx, key).Result()

	return err

}
