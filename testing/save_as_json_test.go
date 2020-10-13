package rediswrapper_test

import (
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

func TestSaveAsJson(t *testing.T) {
	rdb := rediswrapper.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	sampleJsonData := SampleJsonData{}
	err := faker.FakeData(&sampleJsonData)
	if err != nil {
		t.Errorf("sample data build error %s", err.Error())
	}

	key := faker.ID
	err = rdb.SaveAsJson(key, sampleJsonData)
	if err != nil {
		t.Errorf("SaveAsJson failed %s", err.Error())
	}
}
