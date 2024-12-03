package pkg

import (
	"context"
	"fmt"
	"getCurs/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func NewPostgresClient(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PostgreSQL DSN: %w", err)
	}

	// Настройка пула соединений
	poolConfig.MaxConns = int32(cfg.Postgres.PoolMax)
	poolConfig.MinConns = int32(cfg.Postgres.MinCons)
	poolConfig.HealthCheckPeriod = cfg.Postgres.HealthCheckPeriod

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")
	return pool, nil
}

func ClosePostgres(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
		log.Println("PostgreSQL connection pool closed")
	}
}
