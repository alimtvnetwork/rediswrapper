package tests

import (
	"github.com/bxcodec/faker/v3"
	"github.com/go-redis/redis/v8"
)

type SampleJsonData struct {
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
}

var redisClientOptions = &redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
}

var invalidRedisClientOptions = &redis.Options{
	Addr:     "localhost:6399",
	Password: faker.Password(), // no password set
	DB:       10,  // use default DB
}
