package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	"net"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ABFGServer struct {
	Cfg            *config.Config
	Logger         *zap.Logger
	StorageService *services.StorageService
}

func NewServer(cfg *config.Config, log *zap.Logger, ss *services.StorageService) *ABFGServer {
	return &ABFGServer{
		Cfg:            cfg,
		Logger:         log,
		StorageService: ss,
	}
}

func (s *ABFGServer) Run() {
	server := grpc.NewServer()
	l, err := net.Listen("tcp", s.Cfg.Host+":"+s.Cfg.Port)
	if err != nil {
		s.Logger.Sugar().Errorf("%s", err)
	}

	abfg.RegisterABFGuardServiceServer(server, s)

	s.Logger.Sugar().Infof("Server initialisation is completed. Server address: %s:%s", s.Cfg.Host, s.Cfg.Port)
	s.Logger.Sugar().Errorf("%s", server.Serve(l))
}
