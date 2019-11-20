package grpc

import (
	"context"
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	req "github.com/omerkaya1/abf-guard/internal/server"
	"google.golang.org/grpc"
	"log"
	"time"

	abfg "github.com/omerkaya1/abf-guard/internal/server/api"
	"github.com/spf13/cobra"
)

var (
	host, port, login, password, ip, mask string
	black                                 bool
)

var (
	// ClientRootCmd .
	ClientRootCmd = &cobra.Command{
		Use:     "grpc-client",
		Short:   "Run GRPC Web Service client",
		Example: "  abf-guard grpc-client create -h",
	}

	authoriseActionCmd = &cobra.Command{
		Use:     "authorise",
		Short:   "Send an authorisation request",
		Run:     authoriseCmdFunc,
		Example: "  abf-guard grpc-client authorise -s 127.0.0.1 -p 6666 -l user_name -p some_password -i 127.0.0.1",
	}

	flashBucketCmd = &cobra.Command{
		Use:     "flash",
		Short:   "Update calendar event",
		Run:     flashBucketCmdFunc,
		Example: "  abf-guard grpc-client flash -l morty -i 127.0.0.1",
	}

	addIPActionCmd = &cobra.Command{
		Use:     "add",
		Short:   "add ip command",
		Run:     addIPCmdFunc,
		Example: "  abf-guard grpc-client add -b true ",
	}

	deleteIPActionCmd = &cobra.Command{
		Use:     "delete",
		Short:   "delete ip command",
		Run:     deleteIPCmdFunc,
		Example: "  abf-guard grpc-client delete -b true ",
	}

	getIPListActionCmd = &cobra.Command{
		Use:     "get",
		Short:   "get ip list command",
		Run:     getIPListCmdFunc,
		Example: "  abf-guard grpc-client delete -b true ",
	}
)

func init() {
	ClientRootCmd.AddCommand(authoriseActionCmd, flashBucketCmd, addIPActionCmd, deleteIPActionCmd, getIPListActionCmd)
	ClientRootCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "-h, --host=127.0.0.1")
	ClientRootCmd.PersistentFlags().StringVarP(&port, "port", "p", "6666", "-p, --port=7777")
	ClientRootCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "-l, --login=morty")
	ClientRootCmd.PersistentFlags().StringVarP(&password, "password", "w", "", "-w, --password=oh_geez")
	ClientRootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "-i, --ip=127.0.0.1")
	ClientRootCmd.PersistentFlags().StringVarP(&mask, "mask", "m", "", "-m, --mask=")
	ClientRootCmd.PersistentFlags().BoolVarP(&black, "blacklist", "b", false, "-b, --blacklist=true")
}

func authoriseCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || password == "" || ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	ok, err := client.Authorisation(ctx, req.PrepareGRPCAuthorisationBody(login, password, ip))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	log.Println(ok)
}

func flashBucketCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	ok, err := client.FlushBucket(ctx, req.PrepareFlushBucketGrpcRequest(login, ip))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	log.Println(ok)
}

func addIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.AddIPToWhitelist(ctx, req.PrepareSubnetGrpcRequest(ip, black))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetOk())
}

func deleteIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.DeleteIPFromBlacklist(ctx, req.PrepareSubnetGrpcRequest(ip, black))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetOk())
}

func getIPListCmdFunc(cmd *cobra.Command, args []string) {
	client := getGRPCClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.GetIPList(ctx, req.PrepareIPListGrpcRequest(black))
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetResult())
}

func getGRPCClient() abfg.ABFGuardServiceClient {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	return abfg.NewABFGuardServiceClient(conn)
}