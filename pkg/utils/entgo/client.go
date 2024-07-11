package entgo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	entSql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

type IClient interface {
	Close() error
}

type Client[T IClient] struct {
	db  T
	drv *entSql.Driver
}

func NewClient[T IClient](db T, drv *entSql.Driver) *Client[T] {
	return &Client[T]{
		db:  db,
		drv: drv,
	}
}

func (c *Client[T]) Client() T {
	return c.db
}

func (c *Client[T]) Driver() *entSql.Driver {
	return c.drv
}

func (c *Client[T]) DB() *sql.DB {
	return c.drv.DB()
}

// Close database connection
func (c *Client[T]) Close() error {
	return c.db.Close()
}

// Query data
func (c *Client[T]) Query(ctx context.Context, query string, args, v any) error {
	return c.Driver().Query(ctx, query, args, v)
}

// Exec query
func (c *Client[T]) Exec(ctx context.Context, query string, args, v any) error {
	return c.Driver().Exec(ctx, query, args, v)
}

// SetConnectionOption set connection configuration
func (c *Client[T]) SetConnectionOption(maxIdleConnections, maxOpenConnections int, connMaxLifetime time.Duration) {
	// The maximum number of idle connections retained in the connection pool
	c.DB().SetMaxIdleConns(maxIdleConnections)
	// The maximum number of connections the connection pool can open at the same time
	c.DB().SetMaxOpenConns(maxOpenConnections)
	// The maximum length of time a connection can be reused
	c.DB().SetConnMaxLifetime(connMaxLifetime)
}

func driverNameToSemConvKeyValue(driverName string) attribute.KeyValue {
	switch driverName {
	case "mariadb":
		return semconv.DBSystemMariaDB
	case "mysql":
		return semconv.DBSystemMySQL
	case "postgresql":
		return semconv.DBSystemPostgreSQL
	case "sqlite":
		return semconv.DBSystemSqlite
	default:
		return semconv.DBSystemKey.String(driverName)
	}
}

// CreateDriver create database driver
func CreateDriver(driverName, dsn string, enableTrace, enableMetrics bool) (*entSql.Driver, error) {
	var db *sql.DB
	var drv *entSql.Driver
	var err error

	if enableTrace {
		// Connect to database
		if db, err = otelsql.Open(driverName, dsn, otelsql.WithAttributes(
			driverNameToSemConvKeyValue(driverName),
		)); err != nil {
			return nil, errors.New(fmt.Sprintf("failed opening connection to db: %v", err))
		}

		drv = entSql.OpenDB(driverName, db)
	} else {
		if drv, err = entSql.Open(driverName, dsn); err != nil {
			return nil, errors.New(fmt.Sprintf("failed opening connection to db: %v", err))
		}

		db = drv.DB()
	}

	// Register DB stats to meter
	if enableMetrics {
		err = otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
			driverNameToSemConvKeyValue(driverName),
		))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed register otel meter: %v", err))
		}
	}

	return drv, nil
}
