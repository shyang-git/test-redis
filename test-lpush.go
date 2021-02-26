package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	REDIS_URL string
)

var ctx = context.Background()
var key string
var value string
var rdb *redis.Client

func main() {
	const appVersion = "0.9.1"
	versionFlag := flag.Bool("v", false, "prints current version")
	keyFlag := flag.String("k", "ncurion", "redis key")
	// destinationPathFlag := flag.String("d", ".", "destination path")
	// rulesetFlag := flag.String("r", "emerging-all.rules", "etopen ruleset")

	flag.Parse()

	if *versionFlag {
		fmt.Println(appVersion)
		os.Exit(0)
	}
	_ = godotenv.Load("settings.env")
	REDIS_URL = os.Getenv("REDIS_URL")

	r := redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     REDIS_URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := r.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	rdb = r
	key = *keyFlag
	value = "test"

	lpushString(*r, *keyFlag, "test")

	c := cron.New()
	// c.AddFunc("* * * * * *", runEverySecond)
	c.AddFunc("@every 1s", runEverySecond)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func runEverySecond() {
	fmt.Println("runEverySecond")
	lpushString(*rdb, key, value+string(time.Now().Local().Format(time.RFC3339)))
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

func lpushString(r redis.Client, k string, value string) {
	// var keys []string
	intCmds := r.LPush(ctx, k, value)
	cmd := intCmds.Val()

	fmt.Println("Cmd: ", cmd)
}
