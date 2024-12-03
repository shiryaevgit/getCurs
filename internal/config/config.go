package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Postgres `yaml:"postgres" yaml:"postgres"`
	}
	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
	}
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	Postgres struct {
		Host              string        `env-required:"true" yaml:"host"`
		Port              int           `env-required:"true" yaml:"port"`
		User              string        `env-required:"true" yaml:"user"`
		Password          string        `env-required:"true" yaml:"password"`
		DBName            string        `env-required:"true" yaml:"dbname"`
		PoolMax           int           `yaml:"pool_max"`
		MinCons           int           `yaml:"min_cons"`
		HealthCheckPeriod time.Duration `yaml:"health_check_period"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./configs/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("configs error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
