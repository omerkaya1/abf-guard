package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/omerkaya1/abf-guard/internal/config"
	"github.com/omerkaya1/abf-guard/internal/db"
	"github.com/omerkaya1/abf-guard/internal/domain"
	logger "github.com/omerkaya1/abf-guard/internal/log"
	"github.com/omerkaya1/abf-guard/internal/server"
	"github.com/spf13/cobra"
)

const serverErrPrefix = "server error"

var cfgPath string

// ServerRootCmd is the main command to start the GRPC server
var ServerRootCmd = &cobra.Command{
	Use:     "grpc-server",
	Short:   "Run GRPC Server for ABF-Guard",
	Example: "  abf-guard grpc-server -c /path/to/config.json",
	Run: func(cmd *cobra.Command, args []string) {
		// Config file path check
		if cfgPath == "" {
			log.Fatal(errFlagsNotSet)
		}
		// Create the root context for the app
		ctx, cancel := context.WithCancel(context.Background())
		// Start a routine that will monitor OS signals
		go monitorSignalChan(cancel)
		// Initialise configuration
		cfg, err := config.InitConfig(cfgPath)
		assertError(serverErrPrefix, err)
		assertConfig(serverErrPrefix, "invalid server configuration", cfg.Server)
		// Initialise project's logger
		l, err := logger.InitLogger(cfg.Server.Level)
		assertError(serverErrPrefix, err)
		// Init DB
		mainDB, err := db.NewPsqlStorage(&cfg.DB)
		assertError(serverErrPrefix, err)
		// Check the DB configuration
		assertConfig(serverErrPrefix, "invalid DB configuration", cfg.DB)
		// Check the Limits configuration
		assertConfig(serverErrPrefix, "invalid limits configuration", cfg.DB)
		// Init settings for the bucket manager
		mgrSettings, err := domain.InitBucketManagerSettings(&cfg.Limits)
		assertError(serverErrPrefix, err)
		// Check the validity of the bucket manager settings
		assertConfig(serverErrPrefix, "invalid bucket manager settings", mgrSettings)
		// Get bucket manager
		manager, err := domain.NewManager(ctx, mgrSettings)
		assertError(serverErrPrefix, err)
		// Init GRPC server
		srv, err := server.NewServer(&cfg.Server, mainDB, manager)
		assertError(serverErrPrefix, err)
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
	signal.Notify(exitChan, os.Interrupt)
	defer close(exitChan)
	// Listen for OS signals
	for range exitChan {
		cancel()
		return
	}
}
