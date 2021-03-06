package cmd

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/omerkaya1/abf-guard/internal/server"
	"github.com/omerkaya1/abf-guard/internal/server/api"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

const clientErrPrefix = "client error"

var (
	host, port, login, password, ip, entity string
	black                                   bool
)

var (
	errFlagsNotSet = errors.New("cli flags are not set")
	errNotOK       = errors.New("ip was not added to the list")
)

var (
	// ClientRootCmd is the sub-root command for the client
	ClientRootCmd = &cobra.Command{
		Use:     "grpc-client",
		Short:   "Run GRPC Web Service client for ABF-Guard",
		Example: "  abf-guard grpc-client -h",
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
		Short:   "Add an IP to a specified list",
		Run:     addIPCmdFunc,
		Example: "  abf-guard grpc-client add -i 10.0.0.1 -b",
	}

	deleteIPActionCmd = &cobra.Command{
		Use:     "delete",
		Short:   "Delete an IP from a specified list",
		Run:     deleteIPCmdFunc,
		Example: "  abf-guard grpc-client delete -i 10.0.0.1 -b",
	}

	getIPListActionCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get a list of IPs from a specified list",
		Run:     getIPListCmdFunc,
		Example: "  abf-guard grpc-client get -b",
	}
)

func init() {
	ClientRootCmd.AddCommand(authoriseActionCmd, flashBucketCmd, addIPActionCmd, deleteIPActionCmd, getIPListActionCmd, purgeBucketCmd)
	ClientRootCmd.PersistentFlags().StringVarP(&host, "host", "s", "127.0.0.1", "host address")
	ClientRootCmd.PersistentFlags().StringVarP(&port, "port", "p", "6666", "host port")
	ClientRootCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "login parameter")
	ClientRootCmd.PersistentFlags().StringVarP(&password, "password", "w", "", "password parameter")
	ClientRootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "ip parameter")
	ClientRootCmd.PersistentFlags().StringVarP(&entity, "entity", "e", "", "bucket name for removal")
	ClientRootCmd.PersistentFlags().BoolVarP(&black, "blacklist", "b", false, "blacklist or whitelist specification")
}

func authoriseCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || password == "" || ip == "" {
		log.Fatalf("%s: %s", clientErrPrefix, errFlagsNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ok, err := client.Authorisation(ctx, server.PrepareGRPCAuthorisationBody(login, password, ip))
	assertError(clientErrPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", clientErrPrefix, errors.New("the authorisation request was declined"))
	}
	log.Println("the request may proceed")
}

func flashBucketCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || ip == "" {
		log.Fatalf("%s: %s", clientErrPrefix, errFlagsNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ok, err := client.FlushBuckets(ctx, server.PrepareFlushBucketsGrpcRequest(login, ip))
	assertError(clientErrPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", clientErrPrefix, errors.New("flush buckets request failed"))
	}
	log.Printf("%s and %s buckets were flushed", login, ip)
}

func purgeBucketCmdFunc(cmd *cobra.Command, args []string) {
	if entity == "" {
		log.Fatalf("%s: %s", clientErrPrefix, errFlagsNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ok, err := client.PurgeBucket(ctx, server.PreparePurgeBucketGrpcRequest(entity))
	assertError(clientErrPrefix, err)

	if !ok.GetOk() {
		log.Fatalf("%s: %s", clientErrPrefix, errors.New("the requested bucket was not removed"))
	}
	log.Printf("%s bucket was successfully removed", entity)
}

func addIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", clientErrPrefix, errFlagsNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.AddIPToWhitelist(ctx, server.PrepareSubnetGrpcRequest(ip, black))
	assertError(clientErrPrefix, err)

	if resp.GetError() != "" {
		log.Fatalf("%s: %s", clientErrPrefix, resp.GetError())
	}
	if !resp.GetOk() {
		log.Fatalf("%s: %s", clientErrPrefix, errNotOK)
	}
	if black {
		log.Printf("%s was added to the blacklist", ip)
	} else {
		log.Printf("%s was added to the whitelist", ip)
	}
}

func deleteIPCmdFunc(cmd *cobra.Command, args []string) {
	if ip == "" {
		log.Fatalf("%s: %s", clientErrPrefix, errFlagsNotSet)
	}
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.DeleteIPFromBlacklist(ctx, server.PrepareSubnetGrpcRequest(ip, black))
	assertError(clientErrPrefix, err)

	if resp.GetError() != "" {
		log.Fatalf("%s: %s", clientErrPrefix, resp.GetError())
	}
	if !resp.GetOk() {
		log.Fatalf("%s: %s", clientErrPrefix, errNotOK)
	}
	if black {
		log.Printf("%s was deleted from the blacklist", ip)
	} else {
		log.Printf("%s was deleted from the whitelist", ip)
	}
}

func getIPListCmdFunc(cmd *cobra.Command, args []string) {
	client := getGRPCClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.GetIPList(ctx, server.PrepareIPListGrpcRequest(black))
	assertError(clientErrPrefix, err)
	if resp.GetError() != "" {
		log.Fatalf("%s: %s", clientErrPrefix, resp.GetError())
	}
	if black {
		log.Printf("subnet addresses in the blacklist: %v", resp.GetResult())
	} else {
		log.Printf("subnet addresses in the whitelist: %v", resp.GetResult())
	}
}

func getGRPCClient() api.ABFGuardClient {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	assertError(clientErrPrefix, err)
	return api.NewABFGuardClient(conn)
}
