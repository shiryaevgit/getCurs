package repo

import (
	"context"
	"getCurs/internal/entity"
	"getCurs/internal/repo/postgers"
)

type PostgresRepoReceiver interface {
	Get(currency string) (entity.Rate, error)
	GetAll() ([]entity.Rate, error)
}

type receiverRepo struct {
	db postgers.Postgres
}

func NewReceiverRepo(db postgers.Postgres) receiverRepo {
	return receiverRepo{db}
}

func (r receiverRepo) GetAllCurs(ctx context.Context) ([]entity.Rate, error) {
	rates, err := r.db.GetAllRates(ctx)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (r receiverRepo) GetCurs(ctx context.Context, currency string) (entity.Rate, error) {
	//cahedRate, err := r.cache.Get(currency)
	//if err == nil {
	//	return cahedRate, nil
	//}
	//
	//rate, err := r.db.GetRate(ctx, currency)
	//if err != nil {
	//	return entity.Rates{}, err
	//}
	//r.cache.Set(currency, rate)
	//
	//return rate, nil
	return entity.Rate{}, nil
}
