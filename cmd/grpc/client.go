package grpc

import (
	"context"
	"log"
	"time"

	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	req "github.com/omerkaya1/abf-guard/internal/grpc"
	api "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

var (
	host, port, login, password, ip, entity string
	black                                   bool
)

var (
	// ClientRootCmd .
	ClientRootCmd = &cobra.Command{
		Use:     "grpc-client",
		Short:   "Run GRPC Web Service client",
		Example: "  abf-guard grpc-client create -h",
	}

	authoriseActionCmd = &cobra.Command{
		Use:     "auth",
		Short:   "Authorisation request",
		Run:     authoriseCmdFunc,
		Example: "  abf-guard grpc-client auth -s 127.0.0.1 -p 6666 -l user_name -w some_password -i 127.0.0.1",
	}

	flashBucketCmd = &cobra.Command{
		Use:     "flush",
		Short:   "Send a flush buckets request",
		Run:     flashBucketCmdFunc,
		Example: "  abf-guard grpc-client flush -l morty -i 127.0.0.1",
	}

	purgeBucketCmd = &cobra.Command{
		Use:     "purge",
		Short:   "Purge single bucket",
		Run:     purgeBucketCmdFunc,
		Example: "  abf-guard grpc-client purge -e morty",
	}

	addIPActionCmd = &cobra.Command{
		Use:     "add",
		Short:   "Add ip to a specified list",
		Run:     addIPCmdFunc,
		Example: "  abf-guard grpc-client add -i 10.0.0.1 -b",
	}

	deleteIPActionCmd = &cobra.Command{
		Use:     "delete",
		Short:   "Delete ip from a specified list",
		Run:     deleteIPCmdFunc,
		Example: "  abf-guard grpc-client delete -i 10.0.0.1 -b",
	}

	getIPListActionCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get a list of ip from a specified list",
		Run:     getIPListCmdFunc,
		Example: "  abf-guard grpc-client get -b",
	}
)

func init() {
	ClientRootCmd.AddCommand(authoriseActionCmd, flashBucketCmd, addIPActionCmd, deleteIPActionCmd, getIPListActionCmd, purgeBucketCmd)
	ClientRootCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "-h, --host=127.0.0.1")
	ClientRootCmd.PersistentFlags().StringVarP(&port, "port", "p", "6666", "-p, --port=7777")
	ClientRootCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "-l, --login=morty")
	ClientRootCmd.PersistentFlags().StringVarP(&password, "password", "w", "", "-w, --password=oh_geez")
	ClientRootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "-i, --ip=127.0.0.1")
	ClientRootCmd.PersistentFlags().StringVarP(&entity, "entity", "e", "", "-e, --entity=bucket_name")
	ClientRootCmd.PersistentFlags().BoolVarP(&black, "blacklist", "b", false, "-b, --blacklist=true")
}

func authoriseCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || password == "" || ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	ok, err := client.Authorisation(ctx, req.PrepareGRPCAuthorisationBody(login, password, ip))
	oops(errors.ErrClientCmdPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrAuthorisationFailed)
	}
	log.Println("the request may proceed")
	cancel()
}

func flashBucketCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	ok, err := client.FlushBuckets(ctx, req.PrepareFlushBucketsGrpcRequest(login, ip))
	oops(errors.ErrClientCmdPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrFlushBucketsFailed)
	}
	log.Println("the flush request succeeded")
	cancel()
}

func purgeBucketCmdFunc(cmd *cobra.Command, args []string) {
	if entity == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	ok, err := client.PurgeBucket(ctx, req.PreparePurgeBucketGrpcRequest(entity))
	oops(errors.ErrClientCmdPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrPurgeBucketFailed)
	}
	cancel()
	log.Println("the bucket was successfully removed")
}

func addIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.AddIPToWhitelist(ctx, req.PrepareSubnetGrpcRequest(ip, black))
	oops(errors.ErrClientCmdPrefix, err)

	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	if !resp.GetOk() {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrAddIPFailure)
	}
	if black {
		log.Println("ip was added to the blacklist")
	} else {
		log.Println("ip was added to the whitelist")
	}
	cancel()
}

func deleteIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.DeleteIPFromBlacklist(ctx, req.PrepareSubnetGrpcRequest(ip, black))
	oops(errors.ErrClientCmdPrefix, err)

	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	if !resp.GetOk() {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrDeleteIPFailure)
	}
	if black {
		log.Println("ip was deleted from the blacklist")
	} else {
		log.Println("ip was deleted from the whitelist")
	}
	cancel()
}

func getIPListCmdFunc(cmd *cobra.Command, args []string) {
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	resp, err := client.GetIPList(ctx, req.PrepareIPListGrpcRequest(black))
	oops(errors.ErrClientCmdPrefix, err)
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, resp.GetError())
	}
	log.Println(resp.GetResult())
	cancel()
}

func getGRPCClient() api.ABFGuardClient {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	oops(errors.ErrClientCmdPrefix, err)
	return api.NewABFGuardClient(conn)
}
