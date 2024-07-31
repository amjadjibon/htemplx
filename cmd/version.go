package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "htemplx version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("htemplx version: %s\n", VERSION)
		fmt.Printf("build with: %s\n", runtime.Version())
	},
}
