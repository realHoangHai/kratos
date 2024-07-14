package data

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent"
	"github.com/realHoangHai/kratos/app/core/internal/data/ent/migrate"
	"github.com/realHoangHai/kratos/pkg/utils/entgo"
)

// NewEntClient Create Ent ORM database client
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.Client[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "core/data/ent"))

	drv, err := entgo.CreateDriver(
		cfg.Data.Database.GetDriver(),
		cfg.Data.Database.GetSource(),
		cfg.Data.Database.GetEnableTrace(),
		cfg.Data.Database.GetEnableMetrics(),
	)
	if err != nil {
		l.Fatalf("Failed opening connection to db: %v", err)
		return nil
	}

	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)

	// Run the database migration tool
	if cfg.Data.Database.GetMigrate() {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("Failed creating schema resources: %v", err)
		}
	}

	cli := entgo.NewClient(client, drv)

	cli.SetConnectionOption(
		int(cfg.Data.Database.GetMaxIdleConnections()),
		int(cfg.Data.Database.GetMaxOpenConnections()),
		cfg.Data.Database.GetConnectionMaxLifetime().AsDuration(),
	)

	return cli
}
