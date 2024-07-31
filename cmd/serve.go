package cmd

import (
	"htemplx/app/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}
