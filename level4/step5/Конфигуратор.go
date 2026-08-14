package config

type Config struct {
	DefaultBalance float64
}

func NewBalance() *Config {
	return &Config{DefaultBalance: 1000}
}
