package grpc

import (
	"context"
	"net"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/bucket"
	"github.com/omerkaya1/abf-guard/internal/domain/interfaces/db"
	"github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/omerkaya1/abf-guard/internal/log"
	"google.golang.org/grpc"
)

// ABFGuardServer is the object that represents a server for the ABFGuard service
type ABFGuardServer struct {
	// TODO: come up with a way of hiding the configuration behind the interface so that we don't need to pass
	// 		 the configuration by address
	Cfg           *config.Server
	Storage       db.Storage
	BucketManager bucket.Manager
}

// NewServer creates a new ABFGuardServer object and returns it to the callee
func NewServer(cfg *config.Server, sp db.Storage, bm bucket.Manager) (*ABFGuardServer, error) {
	return &ABFGuardServer{
		Cfg:           cfg,
		Storage:       sp,
		BucketManager: bm,
	}, nil
}

// Run starts the ABFGuard server
func (s *ABFGuardServer) Run(ctx context.Context, logger log.Logger) {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		logger.Fatal(err)
	}

	api.RegisterABFGuardServer(server, s)

	// Log errors that occur during the work with buckets
	go func(errChan chan error) {
		for {
			select {
			case <-ctx.Done():
				logger.Info("Context interrupt received")
				if ctx.Err() != nil && ctx.Err() != context.Canceled {
					logger.Infof("Context error: %s", ctx.Err())
				}
				server.GracefulStop()
				logger.Info("Graceful shutdown performed. Bye!")
				return
			case err := <-errChan:
				if err != nil {
					logger.Errorf("%s: %s", errors.ErrBucketManagerPrefix, err)
				}
			}
		}
	}(s.BucketManager.GetErrChan())

	logger.Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	if err = server.Serve(l); err != nil {
		logger.Fatal(err)
	}
}
