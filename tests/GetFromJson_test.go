package tests

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/evatix-go/rediswrapper"
	"testing"
)

func TestGetFromJson(t *testing.T) {
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

	sampleJsonDataReturned := SampleJsonData{}
	err = rdb.GetFromJson(key, &sampleJsonDataReturned)

	if err != nil {
		t.Errorf("GetFromJson failed %s", err.Error())
	}

	if sampleJsonData != sampleJsonDataReturned {
		t.Errorf("GetFromJson didn't return the same data for the key %s", key)
	}
}