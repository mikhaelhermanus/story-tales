package adapter

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	dbredis *redis.Client
)

// LoadRedis is load connection to redis server
func LoadRedis(url, port string) {
	dbredis = Redis(url, port)
}

// Redis is open connection to redis server
func Redis(host, port string) *redis.Client {
	addr := fmt.Sprintf("%v:%v", host, port)
	dbnum, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	opt := redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbnum,
	}
	client := redis.NewClient(&opt)

	return client
}

// UseRedis is open connection into database
func UseRedis() *redis.Client {
	return dbredis
}
