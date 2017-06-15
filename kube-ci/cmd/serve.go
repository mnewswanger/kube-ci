package cmd

import (
	"github.com/spf13/cobra"

	"go.mikenewswanger.com/kube-ci/kube-ci/apiserver"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiserver.StartWebserver(commandLineFlags.datastore, commandLineFlags.listenPort, uint8(commandLineFlags.verbosity))
	},
}

func init() {
	serveCmd.Flags().StringVarP(&commandLineFlags.datastore, "datastore", "d", "", "Datastore Configuration String")
	serveCmd.Flags().Uint16VarP(&commandLineFlags.listenPort, "listen-port", "p", 8080, "Listen Port")

	RootCmd.AddCommand(serveCmd)
}
