package postgres

import (
	"context"

	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pgPool *pgxpool.Pool

type Config struct {
	DatabaseURI string
	MaxConns    int32
	MinConns    int32
}

func Init(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	if pgPool != nil {
		return pgPool, nil
	}
	
	poolCfg, err := pgxpool.ParseConfig(cfg.DatabaseURI)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database uri: %w", err)
	}

	poolCfg.MaxConns = cfg.MaxConns
	poolCfg.MinConns = cfg.MinConns
	poolCfg.MaxConnLifetime = time.Hour
	poolCfg.MaxConnIdleTime = 30 * time.Minute
	poolCfg.HealthCheckPeriod = time.Minute

	pgPool, err = pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	if err := pgPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return pgPool, nil
}

func Get() *pgxpool.Pool {
	if pgPool == nil {
		panic("postgres pool not initialized")
	}
	return pgPool
}

func Close() {
	if pgPool != nil {
		pgPool.Close()
	}
}