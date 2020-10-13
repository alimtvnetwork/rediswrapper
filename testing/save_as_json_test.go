package rediswrapper_test

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/evatix-go/rediswrapper"
	"github.com/go-redis/redis/v8"
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

func TestSaveAsJson(t *testing.T) {
	rdb := rediswrapper.NewClient(context.Background() ,redisClientOptions)

	sampleJsonData := SampleJsonData{}
	err := faker.FakeData(&sampleJsonData)
	if err != nil {
		t.Errorf("sample data build error %s", err.Error())
	}

	key := faker.Name()
	fmt.Println(key, sampleJsonData)
	err = rdb.SaveAsJson(key, sampleJsonData)
	if err != nil {
		t.Errorf("SaveAsJson failed %s", err.Error())
	}
}
