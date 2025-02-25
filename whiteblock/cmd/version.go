package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	// VERSION is set during build
	VERSION = "DEFAULT_VERSION"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get whiteblock CLI client version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " " + VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
