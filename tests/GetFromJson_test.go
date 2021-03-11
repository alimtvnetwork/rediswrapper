package tests

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/evatix-go/rediswrapper"
)

func TestGetFromJson(t *testing.T) {
	rdb := rediswrapper.New(context.Background(), redisClientOptions)

	sampleJsonData := SampleJsonData{}
	err := faker.FakeData(&sampleJsonData)
	if err != nil {
		t.Errorf("sample data build error %s", err.Error())
	}

	key := faker.Name()
	err2 := rdb.SaveAsMarshal(key, sampleJsonData)
	err2.HandleError()

	sampleJsonDataReturned := SampleJsonData{}
	err3 := rdb.GetAsUnmarshal(key, &sampleJsonDataReturned)
	err3.HandleError()
}
