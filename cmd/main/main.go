package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evatix-go/rediswrapper"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

var redisClientOptions = &redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
}

func main()  {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("redis ping error", err.Error())
	} else {
		fmt.Println("ping result", pong)
	}

	err = rdb.Set(ctx, "fruit", "Apple", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// SET key value EX 10 NX
	set, err := rdb.SetNX(ctx, "volatile_key", 1234, 2*time.Second).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("volatile_key stored", set)
	}

	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")

	val, err = rdb.Get(ctx, "volatile_key").Result()
	if err == redis.Nil {
		fmt.Println("volatile_key doesn't exist")
	} else {
		fmt.Println("volatile_key", val)
	}

	// SET key value keepttl NX
	//rdb.SetNX(ctx, "list", []string{"X", "d", "a"}, redis.KeepTTL).Result()
	t , _ := json.Marshal([]string{"X", "d", "a"})
	_, err = rdb.Set(ctx, "list", string(t), 0).Result()
	if err != nil {
		log.Println("listValues set error", err.Error())
	} else {
		listValues, err := rdb.Get(ctx, "list").Result()
		if err != nil {
			log.Println("listValues error", err.Error())
		} else {
			log.Println("listValues", listValues)
		}


		// SORT list LIMIT 0 2 ASC
		rdb.Del(ctx, "plist")
		rdb.LPush(ctx, "plist", 60)
		rdb.LPush(ctx, "plist", 40)
		rdb.LPush(ctx, "plist", 5)

		vals, err := rdb.Sort(ctx, "plist", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
		if err != nil {
			log.Println("redis sort error", err.Error())
		} else {
			log.Println("sorted result", vals)
		}
	}

	rdb.Del(ctx, "list")
	rdb.LPush(ctx, "list", 60)
	rdb.LPush(ctx, "list", 40)
	rdb.LPush(ctx, "list", 5)

	// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
	vals, err := rdb.ZRangeByScoreWithScores(ctx, "list", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
		Offset: 0,
		Count: 2,
	}).Result()

	if err != nil {
		log.Println("ZRangeByScoreWithScores error", err.Error())
	} else {
		log.Println("ZRangeByScoreWithScores result", vals)
	}

	// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
	res, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
		Keys: []string{"zset1", "zset2"},
		Weights: []float64{2, 3},
	}).Result()

	if err != nil {
		log.Println("ZInterStore error", err.Error())
	} else {
		log.Println("ZInterStore result", res)
	}

	// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
	evalResults, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"tey"}, "trello").Result()

	if err != nil {
		log.Println("Eval error", err.Error())
	} else {
		log.Println("Eval result", evalResults)
	}


	// custom command
	doRes, err := rdb.Do(ctx, "set", "var1", "value1").Result()

	if err != nil {
		log.Println("doRes error", err.Error())
	} else {
		log.Println("doRes result", doRes)
	}

	var1Val, err := rdb.Get(ctx, "var1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("var1", var1Val)

	rdbWrapper := rediswrapper.NewClient(context.Background() ,redisClientOptions)
	err = rdbWrapper.SaveBytes("byte_key", []byte("test bytes"))
	if err != nil {
		log.Println("SaveBytes error", err.Error())
	} else {
		log.Println("Bytes were saved.")
	}

	byteData, err := rdbWrapper.GetFromBytes("byte_key")
	if err != nil {
		log.Println("GetFromBytes error", err.Error())
	} else {
		log.Println("GetFromBytes returned", string(byteData))
	}
}