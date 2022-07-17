package redis

import (
	"errors"
	"fmt"
	"invitations-mechanism/config"

	"github.com/go-redis/redis"
)

func RedisConnect() (rdc *redis.Client, err error) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetRedisHost(), config.GetRedisPort()),
		Password: config.GetRedisPassword(),
		DB:       config.GetRedisDB(),
	})

	_, err = rdc.Ping().Result()
	if rdc == nil || err != nil {
		err = errors.New("can't connect to redis client!")
		return
	}

	return
}
