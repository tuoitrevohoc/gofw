package gofw

import (
	"context"
	"log"
	"sync"

	"github.com/DataDog/datadog-go/v5/statsd"
	"go.uber.org/zap"
)

type Service interface {
	Run(ctx context.Context) error
}

type ServiceManager struct {
	services   []Service
	logger     *zap.Logger
	statsd     *statsd.Client
	serviceCtx context.Context
	wg         sync.WaitGroup
}

func NewServiceManager() *ServiceManager {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	statsd, err := statsd.New("statsd_exporter:9125")
	if err != nil {
		logger.Error("Failed to create statsd client", zap.Error(err))
	}

	return &ServiceManager{
		logger:     logger,
		statsd:     statsd,
		serviceCtx: context.Background(),
	}
}

func (s *ServiceManager) Logger() *zap.Logger {
	return s.logger
}

func (s *ServiceManager) Statsd() *statsd.Client {
	return s.statsd
}

func (s *ServiceManager) AddService(service Service) {
	s.services = append(s.services, service)
}

func (s *ServiceManager) AddHealthServer(port string) {
	healthServer := NewHealthServer(s, port)
	s.AddService(healthServer)
}

func (s *ServiceManager) Run() error {
	errChan := make(chan error, len(s.services))

	for _, service := range s.services {
		s.wg.Add(1)
		go func(svc Service) {
			defer s.wg.Done()
			if err := svc.Run(s.serviceCtx); err != nil {
				errChan <- err
			}
		}(service)
	}

	// Wait for all services to complete
	go func() {
		s.wg.Wait()
		close(errChan)
	}()

	// Return the first error encountered, if any
	for err := range errChan {
		return err
	}

	return nil
}
