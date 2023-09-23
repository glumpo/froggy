package config

import (
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Logger Logger   `toml:"logger"`
	API    Telegram `toml:"telegram"`
}

func (c Config) String() string {
	return fmt.Sprintf("Logger: %+v, API: {%s}", c.Logger, c.API)
}

type Logger struct {
	Debug bool `toml:"debug"`
}

type Telegram struct {
	ID   string
	Hash string
}

func (t Telegram) String() string {
	const limit = 4
	id := "too small"
	hash := "too small"
	if len(t.ID) > limit {
		id = t.ID[:limit]
	}
	if len(t.Hash) > limit {
		hash = t.Hash[:limit]
	}
	return fmt.Sprintf("ID: %s, Hash: %s", id, hash)
}

func DefaultCfg() *Config {
	cfg := new(Config)
	cfg.Logger = Logger{
		Debug: true,
	}
	cfg.API = Telegram{
		ID:   "none",
		Hash: "none",
	}

	return cfg
}

var ErrCfgInvalid = fmt.Errorf("invalid config")

func cfgErr(msg string, args ...interface{}) error {
	return fmt.Errorf("%w: %s", ErrCfgInvalid, fmt.Sprintf(msg, args...))
}

func Validate(cfg *Config) error {
	if cfg == nil {
		return cfgErr("nil cfg")
	}

	return nil
}

func UnmarshalToml(data []byte) (cfg *Config, err error) {
	cfg = new(Config)
	err = toml.Unmarshal(data, cfg)
	return
}

func MarshalToml(cfg *Config) ([]byte, error) {
	return toml.Marshal(cfg)
}
