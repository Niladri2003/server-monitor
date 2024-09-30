package cache

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func RedisConnection() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	//fmt.Println("Redis Connection URL: ", redisConnURL)
	if err != nil {
		return nil, err
	}
	//Set Redis options
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}
	fmt.Println(options)
	return redis.NewClient(options), nil
}
