package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/db"
	"github.com/omerkaya1/abf-guard/internal/domain/services"
	"log"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/server"
	logger "github.com/omerkaya1/abf-guard/log"
	"github.com/spf13/cobra"
)

var cfgPath string

// ServerRootCmd is the main command to start the GRPC server
var ServerRootCmd = &cobra.Command{
	Use:     "grpc-server",
	Short:   "Run GRPC Server",
	Example: "  abf-guard grpc-server -c /path/to/config.json",
	Run: func(cmd *cobra.Command, args []string) {
		// Config file path check
		if cfgPath == "" {
			log.Fatalf("%s: %s", errors.ErrServiceCmdPrefix, errors.ErrCLIFlagsAreNotSet)
		}
		// Initialise configuration
		cfg, err := config.InitConfig(cfgPath)
		if err != nil {
			log.Fatalf("%s: InitConfig failed: %s", errors.ErrServiceCmdPrefix, err)
		}
		// Initialise project's logger
		l, err := logger.InitLogger(cfg.Server.Level)
		if err != nil {
			log.Fatalf("%s: InitLogger failed: %s", errors.ErrServiceCmdPrefix, err)
		}
		//Init DB
		mainDB, err := db.NewPsqlStorage(cfg.DB)
		if err != nil {
			log.Fatalf("%s: dbFromConfig failed: %s", errors.ErrServiceCmdPrefix, err)
		}
		//Init GRPC server
		srv := server.NewServer(cfg, l, &services.Storage{Processor: mainDB})
		// Run the GRPC server
		srv.Run()
	},
}

func init() {
	ServerRootCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "-c, --config=/path/to/config.json")
}
