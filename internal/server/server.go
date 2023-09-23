package server

import (
	"context"

	"github.com/glumpo/froggy/internal/model/config"
	"github.com/glumpo/froggy/internal/model/log"
	model "github.com/glumpo/froggy/internal/model/server"
)

type Server struct {
	logger log.Logger
	done   chan struct{}
	ctx    context.Context
	cancel context.CancelFunc

	cfg *config.Config
}

func New(cfg *config.Config, logger log.Logger) model.Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		logger: logger,
		done:   make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}
}

func (s Server) Start() {
	s.logger.Info("Starting")
	go s.run()
	s.logger.Info("Started")
}

func (s Server) Stop() {
	s.logger.Info("Stopping")
	s.cancel()
	<-s.done
	s.logger.Info("Stopped")
}

func (s Server) run() {
	s.logger.Infof("Config: %+w", s.cfg)
	<-s.ctx.Done()
	close(s.done)
}
