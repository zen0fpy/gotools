package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"log"
)

const (
	confFile = `H:\zen0fpy\gotools\demo\redis_demo\config.yml`
)

type Config struct {
	Redis redis.RedisConf
}

func main() {

	var c Config
	conf.MustLoad(confFile, &c)

	client := redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass)

	ok := client.Ping()
	if !ok {
		log.Fatal("failed to connect redis server.")

	}

	var err error

	// set
	key1 := "key1"
	key2 := "key2"

	err = client.Set(key1, "abcedf")
	checkErr(err)
	err = client.Set(key2, "abced")
	checkErr(err)

	// get
	s, err := client.Get(key1)
	checkErr(err)
	fmt.Printf("value: %s\n", s)
	s, err = client.Get(key2)
	checkErr(err)
	fmt.Printf("value: %s\n", s)

	client.Set("dddd", "eeee")
	client.Del("dddd")

	b, _ := client.Exists(key1)
	fmt.Printf("%v\n", b)

	client.Set("ddd", "ffff")
	client.Expire("ddd", 200)

	// bitcount
	val, err := client.BitCount(key1, 1, 64)
	checkErr(err)
	fmt.Printf("bit count: %d\n", val)
	val2, err := client.BitCount(key2, 1, 64)
	checkErr(err)
	fmt.Printf("bit count: %d\n", val2)

	//bitOpAnd
	var andValue string
	result, _ := client.BitOpAnd(andValue, key1, key2)
	fmt.Printf("binOpAnd: %d\n", result)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)

	}
}
