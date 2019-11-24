package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"net"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ABFGuardServer .
type ABFGuardServer struct {
	Cfg            config.Server
	Logger         *zap.Logger
	StorageService *services.Storage
	BucketService  *services.Bucket
}

// NewServer .
func NewServer(cfg config.Server, log *zap.Logger, ss *services.Storage, bs *services.Bucket) *ABFGuardServer {
	return &ABFGuardServer{
		Cfg:            cfg,
		Logger:         log,
		StorageService: ss,
		BucketService:  bs,
	}
}

// Run .
func (s *ABFGuardServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	api.RegisterABFGuardServer(server, s)
	// Log errors that occur during the work with buckets
	go func() {
		for err := range s.BucketService.MonitorErrors() {
			if err != nil {
				s.Logger.Sugar().Errorf("%s: %s", errors.ErrBucketManagerPrefix, err)
			}
		}
	}()

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}
