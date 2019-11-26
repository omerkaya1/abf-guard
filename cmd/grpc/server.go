package grpc

import (
	"log"

	"github.com/omerkaya1/abf-guard/internal/db"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/manager"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket/settings"
	"github.com/omerkaya1/abf-guard/internal/domain/services"

	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/grpc"
	logger "github.com/omerkaya1/abf-guard/log"
	"github.com/spf13/cobra"
)

var cfgPath string

// ServerRootCmd is the main command to start the GRPC server
var ServerRootCmd = &cobra.Command{
	Use:     "grpc-server",
	Short:   "Run GRPC Server for ABF-Guard",
	Example: "  abf-guard grpc-server -c /path/to/config.json",
	Run: func(cmd *cobra.Command, args []string) {
		// Config file path check
		if cfgPath == "" {
			log.Fatalf("%s: %s", errors.ErrServiceCmdPrefix, errors.ErrCLIFlagsAreNotSet)
		}
		// Initialise configuration
		cfg, err := config.InitConfig(cfgPath)
		oops(errors.ErrServiceCmdPrefix, err)
		// Initialise project's logger
		l, err := logger.InitLogger(cfg.Server.Level)
		oops(errors.ErrServiceCmdPrefix, err)
		// Init DB
		mainDB, err := db.NewPsqlStorage(cfg.DB)
		oops(errors.ErrServiceCmdPrefix, err)
		// Init settings for the bucket manager
		mgrSettings, err := settings.InitBucketManagerSettings(cfg.Limits)
		oops(errors.ErrServiceCmdPrefix, err)
		// Init BucketService
		manager, err := manager.NewManager(mgrSettings)
		oops(errors.ErrServiceCmdPrefix, err)
		//Init GRPC server
		srv, err := grpc.NewServer(&cfg.Server, l, &services.Storage{Processor: mainDB}, &services.Bucket{Manager: manager})
		oops(errors.ErrServiceCmdPrefix, err)
		// Run the GRPC server
		srv.Run()
	},
}

func init() {
	ServerRootCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "-c, --config=/path/to/config.json")
}

func oops(prefix string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", prefix, err)
	}
}
