package tests

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/evatix-go/rediswrapper"
	"strings"
	"testing"
)

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
