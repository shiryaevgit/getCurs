package main

import (
	"context"
	"fmt"
	"getCurs/internal/config"
	"getCurs/internal/entity"
	"getCurs/internal/pkg"
	"getCurs/internal/repo"
	"getCurs/internal/repo/postgers"
	"getCurs/internal/usecase"
	"log"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pkg.NewPostgresClient(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Postgres: %v", err)
	}
	defer pkg.ClosePostgres(pool)

	postgresRepo := postgers.NewPostgresRepo(pool)

	receiverRepo := repo.NewReceiverRepo(postgresRepo)
	updaterRepo := repo.NewUpdaterRepo(postgresRepo)

	receiver := usecase.NewReceiver(receiverRepo)
	updater := usecase.NewUpdater(updaterRepo)

	// тесты на запрос
	rates, err := receiver.GetAllRates(ctx)
	if err != nil {
		return
	}
	fmt.Println(rates)

	// тест на сохзранение
	ratess := []entity.Rate{
		{
			Currency: "a",
			Value:    10.01,
			Time:     time.Time{},
		},
		{
			Currency: "b",
			Value:    20.21,
			Time:     time.Time{},
		},
	}

	err = updater.Update(ctx, ratess)
	if err != nil {
		fmt.Printf("Error updating rates: %v\n", err)
	}

}
