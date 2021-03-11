package tests

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/evatix-go/rediswrapper"
)

func TestSaveAsJson(t *testing.T) {
	rdb := rediswrapper.New(context.Background(), redisClientOptions)

	sampleJsonData := SampleJsonData{}
	err := faker.FakeData(&sampleJsonData)
	if err != nil {
		t.Errorf("sample data build error %s", err.Error())
	}

	key := faker.Name()
	err2 := rdb.SaveAsMarshal(key, sampleJsonData)
	err2.HasError()

	rdbInvalid := rediswrapper.New(context.Background(), invalidRedisClientOptions)
	err3 := rdbInvalid.SaveAsMarshal(key, sampleJsonData)
	err3.HandleError()
}
