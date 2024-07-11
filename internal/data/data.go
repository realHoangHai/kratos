package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/realHoangHai/kratos/pkg/utils/entgo"
	"github.com/redis/go-redis/v9"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/realHoangHai/kratos/internal/conf"
	"github.com/realHoangHai/kratos/internal/data/ent"
	rd "github.com/realHoangHai/kratos/pkg/cache/redis"
)

var ProviderSet = wire.NewSet(
	NewData,

	NewEntClient,
	NewRedisClient,

	NewUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	db  *entgo.Client[*ent.Client]
	rdb *redis.Client
}

// NewData .
func NewData(entClient *entgo.Client[*ent.Client], redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		_ = d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient create redis client
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	return rd.NewClient(cfg.Data)
}
