package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/whiteblock/cli/whiteblock/util"
	"io/ioutil"
)

var testCmd = &cobra.Command{
	// Hidden: true,
	Use:   "test <file>",
	Short: "Run test cases.",
	Long: `

This command will read from a file to run a test.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckArguments(cmd, args, 1, 1)

		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			util.PrintErrorFatal(err)
		}

		fmt.Println(util.Prettyp(string(b)))

		var cont map[string]interface{}
		err = json.Unmarshal(b, &cont)
		if err != nil {
			util.PrintErrorFatal(err)
		}

		fmt.Println(cont["build"])
		fmt.Println(cont["netconfig"])
		fmt.Println(cont["rpc"])
		fmt.Println(cont["tests"])

		util.JsonRpcCallAndPrint("add_commands", cont["rpc"])
		res, err := util.JsonRpcCall("build", cont["build"])
		if err != nil {
			util.PrintErrorFatal(err)
		}
		buildListener(res.(string))
		util.JsonRpcCallAndPrint("netem", cont["netconfig"])
		util.JsonRpcCallAndPrint("run_tests", cont["tests"])
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}
