package config

// Config is the general configuration structure.
type Config struct {
	Postgres Postgres `yaml:"postgres"`
	Service  Service  `yaml:"service"`
}

// Postgres represents postgresql db config
type Postgres struct {
	DatabaseURL string `yaml:"database_url"`
}

// Service represents ports configs
type Service struct {
	Port string `yaml:"port"`
}
