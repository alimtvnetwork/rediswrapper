package redisdefaults

import (
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.com/evatix-go/core/constants"

	"github.com/evatix-go/rediswrapper"
)

func NewClientWrapper() *rediswrapper.Wrapper {
	return rediswrapper.New(Context, Options)
}

func NewOptions(
	ip, port string,
) *redis.Options {
	address := ip +
		constants.Colon +
		port

	return &redis.Options{
		Addr: address,
	}
}

func NewOptionsUsingDbPlusIdleTime(
	ip, port string,
	db int,
	idleTime time.Duration,
) *redis.Options {
	address := ip +
		constants.Colon +
		port

	return &redis.Options{
		Addr:               address,
		Dialer:             nil,
		Username:           constants.EmptyString,
		Password:           EmptyPassword,
		DB:                 db,
		MaxRetries:         constants.N10,
		IdleTimeout:        idleTime,
		IdleCheckFrequency: DefaultFrequencyCheck,
	}
}

func NewOptionsUsingAll(
	ip, port, user, password string,
	idleTime time.Duration,
	db int,
) *redis.Options {
	address := ip +
		constants.Colon +
		port

	return &redis.Options{
		Addr:               address,
		Username:           user,
		Password:           password,
		DB:                 db,
		MaxRetries:         constants.N10,
		IdleTimeout:        idleTime,
		IdleCheckFrequency: DefaultFrequencyCheck,
	}
}
