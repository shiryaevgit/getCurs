package postgers

import (
	"context"
	"fmt"
	"getCurs/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(db *pgxpool.Pool) Postgres {
	return Postgres{db}
}

func (r *Postgres) GetAllRates(ctx context.Context) ([]entity.Rate, error) {
	rows, err := r.db.Query(ctx, "SELECT currency, value, time FROM rates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []entity.Rate
	for rows.Next() {
		var rate entity.Rate
		if err := rows.Scan(&rate.Currency, &rate.Value, &rate.Time); err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func (r *Postgres) GetRate(ctx context.Context, currency string) (entity.Rate, error) {
	var rate entity.Rate
	err := r.db.QueryRow(ctx, "SELECT currency, value, time FROM rates WHERE currency=$1", currency).
		Scan(&rate.Currency, &rate.Value, &rate.Time)
	return rate, err
}

func (r *Postgres) Save(ctx context.Context, rates []entity.Rate) error {
	batch := &pgx.Batch{}

	for _, rate := range rates {
		// Вставляем запись без ON CONFLICT
		batch.Queue(`
			INSERT INTO rates (currency, value, time)
			VALUES ($1, $2, $3)
		`, rate.Currency, rate.Value, rate.Time)
	}

	br := r.db.SendBatch(ctx, batch)
	defer br.Close()

	// Проверяем результат выполнения каждого запроса
	for range rates {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("postgers.Save: %w", err)
		}
	}

	return nil
}
