package cache

import "errors"

type Cache interface {
	Retrieve(key string) (interface{}, error)

	Put(key string, value interface{}) error

	Delete(key string) error
}

type RedisCache struct {
	store map[string]interface{}
}

func NewRedisCache() RedisCache {
	return RedisCache{
		store: make(map[string]interface{}),
	}
}

func (r *RedisCache) Retrieve(key string) (interface{}, error) {

	value := r.store[key]

	if value != nil {
		return value, nil
	}

	return nil, errors.New("cannot retrieve from cache.")
}

func (r *RedisCache) Put(key string, value interface{}) error {

	r.store[key] = value

	return nil

}

func (r *RedisCache) Delete(key string) error {

	delete(r.store, key)

	return nil

}
