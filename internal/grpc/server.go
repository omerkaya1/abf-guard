package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	"net"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ABFGServer .
type ABFGServer struct {
	Cfg            *config.Config
	Logger         *zap.Logger
	StorageService *services.Storage
}

// NewServer .
func NewServer(cfg *config.Config, log *zap.Logger, ss *services.Storage) *ABFGServer {
	return &ABFGServer{
		Cfg:            cfg,
		Logger:         log,
		StorageService: ss,
	}
}

// Run .
func (s *ABFGServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Server.Host+":"+s.Cfg.Server.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	abfg.RegisterABFGuardServiceServer(server, s)

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Server.Host, s.Cfg.Server.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}
