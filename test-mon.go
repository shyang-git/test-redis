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
	// periodFlag := flag.Int("p", 1, "monitoring period")
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

	rdb = r
	key = *keyFlag
	value = "test"

	lpushString(*keyFlag, "test")

	c := cron.New()
	// c.AddFunc("* * * * * *", runEverySecond)
	c.AddFunc("@every 1s", runEverySecond)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}

func runEverySecond() {
	keys := getKeys()
	len := getLength()
	ttl := getTTL()
	expire := setExpire(10)
	fmt.Println(string(time.Now().Local().Format(time.RFC3339)), keys, len, ttl, expire)
}

func getLength() int {
	intCmds := rdb.LLen(ctx, key)
	cmd := int(intCmds.Val())

	// fmt.Println("Cmd: ", cmd)
	return cmd
}

func getTTL() int {
	durationCmds := rdb.TTL(ctx, key)
	cmd := int(durationCmds.Val().Seconds())

	// fmt.Println("Cmd: ", cmd)
	return cmd
}

func setExpire(seconds int) bool {
	var expire time.Duration
	expire = time.Duration(seconds * 1000000000)

	boolCmds := rdb.Expire(ctx, key, expire)
	cmd := boolCmds.Val()

	// fmt.Println("Cmd: ", cmd)
	return cmd
}

func getKeys() []string {
	// var keys []string
	stringSliceCmds := rdb.Keys(ctx, "*")
	keys := stringSliceCmds.Val()

	// for stringSliceCmd := range stringSliceCmds.val {
	// 	append(keys, stringSliceCmd.)
	// }
	return keys
}

func lpushString(k string, value string) {
	// var keys []string
	intCmds := rdb.LPush(ctx, k, value)
	cmd := intCmds.Val()

	fmt.Println("Cmd: ", cmd)
}
