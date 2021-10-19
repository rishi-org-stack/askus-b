package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//will figure it out later
var ctx = context.TODO()

// var DB *badger.DB
type RedisDB struct {
	DB *redis.Client
}

func Connect() (*RedisDB, error) {

	return &RedisDB{
		DB: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}, nil
}

func (Bdb *RedisDB) Set(key string, val string, exp time.Duration) error {
	err := Bdb.DB.Set(ctx, key, val, exp).Err()
	if err != nil {
		panic(err)
	}
	return err
}

func (Bdb *RedisDB) Get(key string) (string, error) {
	val, err := Bdb.DB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (Bdb *RedisDB) Delete() error {
	err := Bdb.DB.FlushAll(ctx).Err()
	return err
}
