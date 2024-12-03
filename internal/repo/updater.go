package repo

import (
	"context"
	"fmt"
	"getCurs/internal/entity"
	"getCurs/internal/repo/postgers"
)

type updaterRepo struct {
	db postgers.Postgres
	//какое-то соединение с внешним апи
}

func NewUpdaterRepo(db postgers.Postgres) updaterRepo {
	return updaterRepo{
		db: db,
	}
}

func (u updaterRepo) SaveRate(ctx context.Context, rate []entity.Rate) error {
	err := u.db.Save(ctx, rate)
	if err != nil {
		return fmt.Errorf("repo.SaveRate: %w", err)
	}

	return nil
}
func (u updaterRepo) FetchRates() ([]entity.Rate, error) {
	return []entity.Rate{}, nil
}
