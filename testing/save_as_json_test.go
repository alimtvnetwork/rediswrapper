package rediswrapper_test

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/evatix-go/rediswrapper"
	"github.com/go-redis/redis/v8"
	"strings"
	"testing"
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

func TestSaveAsJson(t *testing.T) {
	rdb := rediswrapper.NewClient(context.Background() ,redisClientOptions)

	sampleJsonData := SampleJsonData{}
	err := faker.FakeData(&sampleJsonData)
	if err != nil {
		t.Errorf("sample data build error %s", err.Error())
	}

	key := faker.Name()
	err = rdb.SaveAsJson(key, sampleJsonData)
	if err != nil {
		t.Errorf("SaveAsJson failed %s", err.Error())
	}

	rdbInvalid := rediswrapper.NewClient(context.Background() ,invalidRedisClientOptions)
	err = rdbInvalid.SaveAsJson(key, sampleJsonData)
	if err != nil {
		if !strings.Contains(err.Error(), "connection refused") {
			t.Error("SaveAsJson failed", err.Error())
		}
	}
}
