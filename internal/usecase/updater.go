package usecase

import (
	"context"
	"fmt"
	"getCurs/internal/entity"
)

type updater struct {
	repo updaterRepo
}

type updaterRepo interface {
	SaveRate(ctx context.Context, rate []entity.Rate) error
}

func NewUpdater(repo updaterRepo) updater {
	return updater{repo}
}

func (u *updater) Update(ctx context.Context, rates []entity.Rate) error {
	// Дополнительная логика для обработки данных перед сохранением
	//Принимаем rates

	err := u.repo.SaveRate(ctx, rates)
	if err != nil {
		return fmt.Errorf("usecase.Update: %w", err)
	}
	return nil
}
func (u *updater) StartAutoUpdate(minutes int) error {
	return nil
}
func (u *updater) StopAutoUpdate() error {
	return nil
}
