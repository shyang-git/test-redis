package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var (
	REDIS_URL string
)

var ctx = context.Background()

func ExampleClient(r *redis.Client) {
	err := r.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := r.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := r.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	const appVersion = "0.9.1"
	versionFlag := flag.Bool("v", false, "prints current version")
	testFlag := flag.Bool("t", false, "connection test")
	// destinationPathFlag := flag.String("d", ".", "destination path")
	// rulesetFlag := flag.String("r", "emerging-all.rules", "etopen ruleset")

	flag.Parse()

	if *versionFlag {
		fmt.Println(appVersion)
		os.Exit(0)
	}
	_ = godotenv.Load("settings.env")
	REDIS_URL = os.Getenv("REDIS_URL")

	rdb := redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     REDIS_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	if *testFlag {
		ExampleClient(rdb)
	}

	keys := getKeys(rdb)
	fmt.Println(keys)

	for i, key := range keys {
		durationCmd := rdb.TTL(ctx, key)
		ttl := durationCmd.Val()
		fmt.Println(i, key, ttl)
	}
}

func getKeys(r *redis.Client) []string {
	// var keys []string
	stringSliceCmds := r.Keys(ctx, "*")
	keys := stringSliceCmds.Val()

	// for stringSliceCmd := range stringSliceCmds.val {
	// 	append(keys, stringSliceCmd.)
	// }
	return keys
}
