package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/glumpo/froggy/internal/log"
	"github.com/glumpo/froggy/internal/model/config"
	serverModel "github.com/glumpo/froggy/internal/model/server"
	"github.com/glumpo/froggy/internal/server"
)

var (
	version = "0.0.0"
	commit  = "none"
	date    = "none"
)

func main() {
	if err := do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal froggy error: %v", err)
		os.Exit(1)
	}
}

func do() error {
	var printVersion bool
	var cfgPath string
	flag.BoolVar(&printVersion, "v", false, "Get version")
	flag.StringVar(&cfgPath, "c", "config.toml", "Provide config path")
	flag.Parse()

	if printVersion {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Date: %s\n", date)
		fmt.Printf("Build Commit: %s\n", commit)
		return nil
	}

	// init
	cfgRaw, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to read config %s: %w", cfgPath, err)
	}
	cfg, err := config.UnmarshalToml(cfgRaw)
	if err != nil {
		return fmt.Errorf("failed to unmarshal toml config: %w", err)
	}
	if err := config.Validate(cfg); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	logger := log.New(cfg.Logger)
	logger.Infof("Version: %s, build data: %s, build commit: %s", version, date, commit)
	logger.Debugf("Config path: %s", cfgPath)
	logger.Debugf("Config: %s", cfg)

	// start
	var s serverModel.Server
	s = server.New(cfg, logger.WithSource("server"))
	err = s.Start()
	if err != nil {
		return fmt.Errorf("starting server: %w", err)
	}

	// shutdown handling
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	logger.Info("Initiated graceful shutdown")
	s.Stop()
	logger.Info("Shutdown finished")

	return nil
}
