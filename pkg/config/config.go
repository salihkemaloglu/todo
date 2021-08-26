package config

// Config is the general configuration structure.
type Config struct {
	Postgres      Postgres      `yaml:"postgres"`
	Service       Service       `yaml:"service"`
	Queue         Queue         `yaml:"queue"`
	OnlineService OnlineService `yaml:"online_service"`
}

type Postgres struct {
	DatabaseURL string `yaml:"database_url"`
}

type Service struct {
	Port string `yaml:"port"`
}

type OnlineService struct {
	URL string `yaml:"url"`
}

type Queue struct {
	URL string `yaml:"url"`
}
