package cmd

import (
	"github.com/spf13/cobra"
	"github.com/whiteblock/cli/whiteblock/util"
)

var killCmd = &cobra.Command{
	Aliases: []string{},
	Use:     "kill <node>",
	Short:   "Raise SIGINT to a node's main process and wait for it to die",
	Long: `Sends SIGINT to the node's main process, and continue to query the state of that 
	process until it dies. 
`,

	Run: func(cmd *cobra.Command, args []string) {
		util.CheckArguments(cmd, args, 1, 1)
		testnetId, err := getPreviousBuildId()
		if err != nil {
			util.PrintErrorFatal(err)
		}
		util.JsonRpcCallAndPrint("kill_node", []interface{}{testnetId, args[0]})
	},
}

func init() {
	RootCmd.AddCommand(killCmd)
}
