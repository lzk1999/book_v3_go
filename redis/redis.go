package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"time"
)

var (
	Rdb *redis.Client
	ctx = context.Background()
)

func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
