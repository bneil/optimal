package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	"github.com/kkyr/fig"
)

type Config struct {
	App struct {
		Environment string `fig:"environment" validate:"required"`
	} `fig:"app"`
	Server struct {
		Host         string        `fig:"host" default:"0.0.0.0"`
		Port         int           `fig:"port" default:"80"`
		ReadTimeout  time.Duration `fig:"read_timeout" default:"30s"`
		WriteTimeout time.Duration `fig:"write_timeout" default:"30s"`
	} `fig:"server"`
	Database struct {
		Location string `fig:"location" default:"database/local.db"`
	} `fig:"database"`
	Logger struct {
		Level   string         `fig:"level" default:"info"`
		Pattern *regexp.Regexp `fig:"pattern" default:".*"`
	} `fig:"logger"`
}

var ConfigPointer *Config

func GetConfig() *Config {
	if ConfigPointer == nil {
		var cfg Config
		_, file, _, _ := runtime.Caller(0)
		rootDir := filepath.Dir(filepath.Dir(file))
		env := findEnvironmentConfig()
		filename := fmt.Sprintf("%s.yaml", env)
		dirs := fmt.Sprintf("%s/config", rootDir)
		err := fig.Load(&cfg,
			fig.File(filename),
			fig.Dirs(dirs),
		)
		if err != nil {
			panic(err)
		}
		ConfigPointer = &cfg
		cfg.Database.Location = fmt.Sprintf("%s/%s", rootDir, cfg.Database.Location)
		return ConfigPointer
	}

	return ConfigPointer
}

func findEnvironmentConfig() string {
	var env string
	switch os.Getenv("ENV") {
	case "production", "prod", "prd":
		env = "prod"
	case "development", "dev":
		env = "dev"
	default:
		env = "local"
	}
	return env
}
