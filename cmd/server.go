package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/omerkaya1/abf-guard/internal/db"
	"github.com/omerkaya1/abf-guard/internal/domain/bucket"
	"github.com/omerkaya1/abf-guard/internal/domain/config"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"github.com/omerkaya1/abf-guard/internal/grpc"
	logger "github.com/omerkaya1/abf-guard/internal/log"
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
			panic(errors.ErrCLIFlagsAreNotSet)
		}
		// Create the root context for the app
		ctx, cancel := context.WithCancel(context.Background())
		// Start a routine that will monitor OS signals
		go monitorSignalChan(cancel)
		// Initialise configuration
		cfg, err := config.InitConfig(cfgPath)
		assertError(errors.ErrServiceCmdPrefix, err)
		assertConfig(errors.ErrServiceCmdPrefix, "invalid server configuration", cfg.Server)
		// Initialise project's logger
		l, err := logger.InitLogger(cfg.Server.Level)
		assertError(errors.ErrServiceCmdPrefix, err)
		// Init DB
		mainDB, err := db.NewPsqlStorage(&cfg.DB)
		assertError(errors.ErrServiceCmdPrefix, err)
		// Check the DB configuration
		assertConfig(errors.ErrServiceCmdPrefix, "invalid DB configuration", cfg.DB)
		// Check the Limits configuration
		assertConfig(errors.ErrServiceCmdPrefix, "invalid limits configuration", cfg.DB)
		// Init settings for the bucket manager
		mgrSettings, err := bucket.InitBucketManagerSettings(&cfg.Limits)
		assertError(errors.ErrServiceCmdPrefix, err)
		// Check the validity of the bucket manager settings
		assertConfig(errors.ErrServiceCmdPrefix, "invalid bucket manager settings", mgrSettings)
		// Get bucket manager
		manager, err := bucket.NewManager(ctx, mgrSettings)
		assertError(errors.ErrServiceCmdPrefix, err)
		// Init GRPC server
		srv, err := grpc.NewServer(&cfg.Server, mainDB, manager)
		assertError(errors.ErrServiceCmdPrefix, err)
		// Run the GRPC server
		srv.Run(ctx, l.Sugar())
	},
}

func init() {
	ServerRootCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "-c, --config=/path/to/config.json")
}

// A shortcut to assert critical errors
func assertError(prefix string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", prefix, err)
	}
}

// A shortcut to assert critical parts of configurations
func assertConfig(prefix, message string, cfg config.Validator) {
	if !cfg.Valid() {
		log.Fatalf("%s: %s", prefix, message)
	}
}

// Monitors the os interruption signals in order to gracefully shut down the service and release all associated resources
func monitorSignalChan(cancel context.CancelFunc) {
	// Handle interrupt
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, os.Kill)
	defer close(exitChan)
	// Listen for OS signals
	for range exitChan {
		cancel()
		return
	}
}
