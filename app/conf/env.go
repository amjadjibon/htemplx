package conf

import (
	"time"

	"github.com/caarlos0/env/v11"
)

//go:generate go run github.com/g4s8/envdoc -output ../../docs/env.md -format markdown
type Config struct {
	HttpHost         string        `env:"HOST" envDefault:"0.0.0.0"`
	HttpPort         int           `env:"PORT" envDefault:"8080"`
	HttpReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"10s"`
	HttpWriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"10s"`
	LogLevel         string        `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() Config {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		panic(err)
	}

	return cfg
}
