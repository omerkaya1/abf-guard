package grpc

import (
	"github.com/omerkaya1/abf-guard/internal/domain/errors"
	"google.golang.org/grpc"

	"log"

	abfg "github.com/omerkaya1/abf-guard/internal/grpc/api"
	"github.com/spf13/cobra"
)

var (
	host, port, login, password, ip, mask string
	black                                 bool
)

var (
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

	addIpActionCmd = &cobra.Command{
		Use:     "add",
		Short:   "add ip command",
		Run:     addIpCmdFunc,
		Example: "  abf-guard grpc-client add -b true ",
	}

	deleteIpActionCmd = &cobra.Command{
		Use:     "delete",
		Short:   "delete ip command",
		Run:     deleteIpCmdFunc,
		Example: "  abf-guard grpc-client delete -b true ",
	}
)

func init() {
	ClientRootCmd.AddCommand(authoriseActionCmd, flashBucketCmd, addIpActionCmd, deleteIpActionCmd)
	ClientRootCmd.Flags().StringVarP(&host, "host", "s", "127.0.0.1", "-h, --host=127.0.0.1")
	ClientRootCmd.Flags().StringVarP(&port, "port", "p", "6666", "-p, --port=7777")
	ClientRootCmd.Flags().StringVarP(&login, "login", "l", "", "-l, --login=morty")
	ClientRootCmd.Flags().StringVarP(&password, "password", "w", "", "-w, --password=oh_jeez")
	ClientRootCmd.Flags().StringVarP(&ip, "ip", "i", "", "-i, --ip=127.0.0.1")
	ClientRootCmd.Flags().StringVarP(&mask, "mask", "m", "", "-m, --mask=")
	ClientRootCmd.Flags().BoolVarP(&black, "blacklist", "b", false, "-b, --blacklist=true")
}

func authoriseCmdFunc(cmd *cobra.Command, args []string) {
	if login == "" || password == "" || ip == "" {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, errors.ErrCLIFlagsAreNotSet)
	}

	//client := getGRPCClient()
	//ctx := context.Background()
	//rb := models.Authorisation{
	//	Login:    login,
	//	Password: password,
	//	IP:       ip,
	//}
	//ok, err := client.Authorisation(ctx, )

}

func flashBucketCmdFunc(cmd *cobra.Command, args []string) {

}

func addIpCmdFunc(cmd *cobra.Command, args []string) {

}

func deleteIpCmdFunc(cmd *cobra.Command, args []string) {

}

func getGRPCClient() abfg.ABFGuardServiceClient {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s: %s", errors.ErrClientCmdPrefix, err)
	}
	return abfg.NewABFGuardServiceClient(conn)
}
