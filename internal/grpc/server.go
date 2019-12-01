package grpc

import (
	"context"
	"net"

	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ABFGuardServer is the object that represents a server for the ABFGuard service
type ABFGuardServer struct {
	Cfg            *config.Server
	Logger         *zap.Logger
	StorageService *services.Storage
	BucketService  *services.Bucket
	ctx            context.Context
}

// NewServer creates a new ABFGuardServer object and returns it to the callee
func NewServer(ctx context.Context, cfg *config.Server, l *zap.Logger, ss *services.Storage, bs *services.Bucket) (*ABFGuardServer, error) {
	if cfg == nil || l == nil || ss == nil || bs == nil {
		return nil, errors.ErrMissingServerParameters
	}
	return &ABFGuardServer{
		ctx:            ctx,
		Cfg:            cfg,
		Logger:         l,
		StorageService: ss,
		BucketService:  bs,
	}, nil
}

// Run starts the ABFGuard server
func (s *ABFGuardServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Fatalf("%s", err)
	}

	api.RegisterABFGuardServer(server, s)

	// Log errors that occur during the work with buckets
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				s.Logger.Sugar().Info("Context interrupt received")
				if s.ctx.Err() != nil && s.ctx.Err() != context.Canceled {
					s.Logger.Sugar().Infof("Context error: %s", s.ctx.Err())
				}
				server.GracefulStop()
				s.Logger.Sugar().Info("Graceful shutdown performed. Bye!")
				return
			case err := <-s.BucketService.MonitorErrors():
				if err != nil {
					s.Logger.Sugar().Errorf("%s: %s", errors.ErrBucketManagerPrefix, err)
				}
			}
		}
	}()

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	if err := server.Serve(l); err != nil {
		s.Logger.Sugar().Fatal(err)
	}
}
