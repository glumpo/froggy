package server

import (
	"context"
	"sync"

	"github.com/glumpo/froggy/internal/model/config"
	"github.com/glumpo/froggy/internal/model/log"
)

type Server struct {
	m       sync.Mutex
	started bool
	ctx     context.Context
	cancel  context.CancelFunc
	done    chan struct{}

	logger log.Logger
	cfg    *config.Config
}

func New(cfg *config.Config, logger log.Logger) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		logger: logger,
		done:   make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}
}

func (s *Server) Start() error {
	s.m.Lock()
	defer s.m.Unlock()
	if s.started {
		s.logger.Debug("Already started")
		return nil
	}
	s.logger.Info("Starting")

	go s.run()

	s.started = true
	s.logger.Info("Started")
	return nil
}

func (s *Server) Stop() {
	s.m.Lock()
	done := s.stop()
	s.m.Unlock()
	if done != nil {
		<-done
		s.logger.Info("Stopped")
	}
}

func (s *Server) stop() chan struct{} {
	if !s.started {
		s.logger.Debug("Not started")
		return nil
	}
	s.started = false

	s.logger.Info("Stopping")
	s.cancel()
	return s.done
}

func (s *Server) run() {
	s.logger.Infof("Current config: %s", s.cfg)
	defer func() {
		s.logger.Info("Exiting")
		close(s.done)
	}()
}
