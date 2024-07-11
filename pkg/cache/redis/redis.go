package redis

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/realHoangHai/kratos/internal/conf"
)

// NewClient create go-redis client
func NewClient(conf *conf.Data) (rdb *redis.Client) {
	if rdb = redis.NewClient(&redis.Options{
		Addr: conf.GetRedis().GetAddr(),
	}); rdb == nil {
		log.Fatalf("Failed opening connection to redis")
		return nil
	}

	return rdb
}
