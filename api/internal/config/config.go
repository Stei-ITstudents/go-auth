package config

import "github.com/caarlos0/env/v9"

type Config struct {
	Port      int    `env:"PORT"       envDefault:"8080"`
	SecretKey string `env:"SECRET_KEY" envDefault:"my-Secret-Key"`
	MySqlDns  string `env:"DB_DNS"     envDefault:"root:password@tcp(localhost:3306)/dbname"`
}

func GetConfigs() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
