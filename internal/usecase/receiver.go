package usecase

import (
	"context"
	"fmt"
	"getCurs/internal/entity"
)

type receiver struct {
	repo receiverRepo
}

type receiverRepo interface {
	GetAllCurs(ctx context.Context) ([]entity.Rate, error)
	GetCurs(ctx context.Context, currency string) (entity.Rate, error)
}

func NewReceiver(repo receiverRepo) receiver {
	return receiver{repo}
}

func (r *receiver) GetAllRates(ctx context.Context) ([]entity.Rate, error) {
	rates, err := r.repo.GetAllCurs(ctx)
	if err != nil {
		return []entity.Rate{}, fmt.Errorf("usecase.GetAllRates: %w", err)
	}
	// здесь может быть дополнительная бизнес логика приложения
	return rates, nil
}
func (r *receiver) GetRateByCurrency(ctx context.Context, currency string) (entity.Rate, error) {
	// здесь может быть дополнительная бизнес логика приложения
	rate, err := r.repo.GetCurs(ctx, currency)
	if err != nil {
		return entity.Rate{}, err
	}
	// здесь может быть дополнительная бизнес логика приложения
	return rate, nil
}
func (r *receiver) CalculateStatistics(ctx context.Context, currency string) (minPrice, maxPrice, hourChange float64, err error) {
	return 0, 0, 0, nil
}
