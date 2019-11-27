package grpc

import (
	"net"
	"os"
	"os/signal"
	"syscall"

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
}

// NewServer creates a new ABFGuardServer object and returns it to the callee
func NewServer(cfg *config.Server, l *zap.Logger, ss *services.Storage, bs *services.Bucket) (*ABFGuardServer, error) {
	if cfg == nil || l == nil || ss == nil || bs == nil {
		return nil, errors.ErrMissingServerParameters
	}
	return &ABFGuardServer{
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

	// Handle interrupt
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)

	// Log errors that occur during the work with buckets
	go func() {
		for {
			select {
			case <-exitChan:
				s.Logger.Sugar().Info("Interrupt signal received.")
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
	s.Logger.Sugar().Infof("%s", server.Serve(l))
}
