package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/whiteblock/cli/whiteblock/util"
	"strconv"
)

var (
	txsFlag      int
	tpsFlag      int
	valueFlag    int
	txSizeFlag   int
	gasFlag      string
	gasPriceFlag string
	toFlag       string
	fromFlag     string
	nodeFlag     string
)

var txCmd = &cobra.Command{
	// Hidden: true,
	Use:   "tx <command>",
	Short: "Run transaction commands.",
	Long: `
The user must specify the blockchain flag as well as any other flags that will be used for sending transactions.
Tx will run commands relevant to sending transactions.

Please use the help commands to make sure you provide the correct flags. If the blockchain is not listed in the help command, the transaction command is not supported for that blockchain. 
	`,
	Run: util.PartialCommand,
}

/*
The primary use of these methods is to be able to send one line commands through a
testing script that will be able to automate transaction tests.
*/

var sendSingleTxCmd = &cobra.Command{
	// Hidden: true,
	Use:   "send",
	Short: "Sends a single transaction",
	Long: `
The user must specify the flags that will be used for sending transactions.
Send a transaction between two accounts.

Required Parameters: 
	ethereum:  --from <address>  --destination <address> --gas <gas> --gasprice <gas price> --value <amount>
	eos:  --node <node> --from <address> --destination <address> --value <amount> 
	
Optional Parameters:
	eos:  --symbol [symbol=SYS] --code [code=eosio.token] --memo [memo=]

`,
	Run: func(cmd *cobra.Command, args []string) {
		command := ""
		params := []string{}

		/*previousBuild, err := getPreviousBuild()
		if err != nil {
			util.PrintErrorFatal(err)
		}*/

		if !(len(toFlag) > 0) || !(len(fromFlag) > 0) || !(len(gasFlag) > 0) || !(len(gasPriceFlag) > 0) || valueFlag == 0 {
			fmt.Println("Required flags were not provided. Please input the required flags.")
			cmd.Help()
			return
		}
		command = "eth::send_transaction"
		params = []string{fromFlag, toFlag, gasFlag, gasPriceFlag, strconv.Itoa(valueFlag)}
		/*
			switch previousBuild.Blockchain {
			case "eos":
				if !(len(nodeFlag) > 0) || !(len(toFlag) > 0) || !(len(fromFlag) > 0) || valueFlag == 0 {
					fmt.Println("Required flags were not provided. Please input the required flags.")
					return
				}
				command = "eos::send_transaction"
				params = []string{nodeFlag, fromFlag, toFlag, strconv.Itoa(valueFlag)}
			default:
				util.ClientNotSupported(previousBuild.Blockchain)
			}
		*/
		util.JsonRpcCallAndPrint(command, params)
	},
}

var startTxCmd = &cobra.Command{
	// Hidden: true,
	Use:   "start",
	Short: "Send transactions",
	Long: `
This command will be used to automate transactions and will require require flags to execute. There are two modes of sending transactions: stream and burst. 
The user must specify the flags that will be used for sending transactions.
	`,
	Run: util.PartialCommand,
}

var startStreamTxCmd = &cobra.Command{
	// Hidden: true,
	Use:     "stream",
	Short:   "Send continuous transactions",
	Aliases: []string{"cont", "continuous"},
	Long: `
The user must specify the blockchain flag as well as any other flags that will be used for sending transactions.
This command will start sending a continual stream of transactions according to the given flags. Stream will send transactions as a continuous flow of tps. The user will need to run the command tx stop to stop running transactions.

Required Parameters: 
	ethereum:  --tps <tps> --value <amount>
	eos:  --tps <tps> 

Optional Parameters:
	ethereum:  --destination [address]
	eos:  --size [tx size]
	`,

	Run: func(cmd *cobra.Command, args []string) {

		if !cmd.Flags().Changed("tps") { //TPS will always be required
			fmt.Println("No \"tps\" flag has been provided. Please input the tps flag with a value.")
			cmd.Help()
			return
		}
		tps, err := cmd.Flags().GetInt("tps")
		if err != nil {
			util.PrintErrorFatal(err)
		}
		params := []interface{}{strconv.Itoa(tps)}
		//value parameter
		value, err := cmd.Flags().GetInt("value")
		if err != nil {
			util.PrintErrorFatal(err)
		}
		toEth := strconv.Itoa(value) + "000000000000000000"
		params = append(params, toEth)

		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			util.PrintErrorFatal(err)
		}
		params = append(params, size)
		dest, err := cmd.Flags().GetString("destination")
		if err != nil {
			util.PrintErrorFatal(err)
		}
		if len(dest) > 0 {
			params = append(params, dest)
		}
		log.WithFields(log.Fields{"params": params}).Debug("Sending the request to start sending tx")
		util.JsonRpcCallAndPrint("run_constant_tps", params)
	},
}

var startBurstTxCmd = &cobra.Command{
	// Hidden: true,
	Use:   "burst",
	Short: "Send burst transactions",
	Long: `
The user must specify the blockchain flag as well as any other flags that will be used for sending transactions.
This command will send a burst of transactions. Additional flags are optional. Burst will send one burst of transactions to the blockchain to fill the transaction pool.

Required Parameters: 
	--txs <number of tx>
	--value <value>
Optional Parameters:
	--size [tx size]
`,
	Run: func(cmd *cobra.Command, args []string) {
		params := []string{strconv.Itoa(txsFlag)}
		previousBuild, err := getPreviousBuild()
		if err != nil {
			util.PrintErrorFatal(err)
		}

		switch previousBuild.Blockchain {
		case "eos":
			//error handling for invalid flags
			if valueFlag != 0 {
				fmt.Println("Invalid \"valueFlag\" flag has been provided.")
				cmd.Help()
				return
			}
			if tpsFlag == 0 {
				fmt.Println("No \"txsFlag\" flag has been provided. Please input the tps flag with a value.")
				cmd.Help()
				return
			}
			if txSizeFlag >= 174 {
				params = append(params, strconv.Itoa(txSizeFlag))
			} else if txSizeFlag > 0 && txSizeFlag < 174 {
				fmt.Println("Transaction size value is too small. The minimum size of a transaction is 174 bytes.")
				return
			}
		default:
			util.ClientNotSupported(previousBuild.Blockchain)
		}
		util.JsonRpcCallAndPrint("run_burst_tx", params)
	},
}

var stopTxCmd = &cobra.Command{
	// Hidden: true,
	Use:   "stop",
	Short: "Stop transactions",
	Long: `
The user must specify the blockchain flag as well as any other flags that will be used for sending transactions.
Stops the sending of transactions if transactions are currently being sent
`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := util.JsonRpcCall("state::kill", []string{})
		if res != nil && res.(float64) == 0 && err == nil {
			fmt.Println("Transactions stopped successfully")
		} else {
			fmt.Println("There was an error stopping transactions")
		}
	},
}

func init() {

	sendSingleTxCmd.Flags().StringVarP(&toFlag, "destination", "d", "", "where the transaction will be sent to")
	sendSingleTxCmd.Flags().StringVarP(&fromFlag, "from", "f", "", "where the transaction will be sent from")
	sendSingleTxCmd.Flags().StringVarP(&gasFlag, "gas", "g", "", "specify gas for tx")
	sendSingleTxCmd.Flags().StringVarP(&nodeFlag, "node", "n", "", "specify node to send tx")
	sendSingleTxCmd.Flags().StringVarP(&gasPriceFlag, "gasprice", "p", "", "specify gas price for tx")
	sendSingleTxCmd.Flags().IntVarP(&valueFlag, "value", "v", 0, "amount to send in transaction")

	startStreamTxCmd.Flags().StringP("destination", "d", "", "where the transaction will be sent to")
	startStreamTxCmd.Flags().IntP("size", "s", 0, "size of the transaction in bytes")
	startStreamTxCmd.Flags().IntP("tps", "t", 0, "transactions per second")
	startStreamTxCmd.Flags().IntP("value", "v", -1, "amount to send in transaction")
	startStreamTxCmd.MarkFlagRequired("tps")

	startBurstTxCmd.Flags().StringVarP(&toFlag, "destination", "d", "", "where the transaction will be sent to")
	startBurstTxCmd.Flags().IntVarP(&txSizeFlag, "size", "s", 0, "size of the transaction in bytes")
	startBurstTxCmd.Flags().IntVarP(&txsFlag, "txs", "t", 0, "transactions per second")
	startBurstTxCmd.Flags().IntVarP(&valueFlag, "value", "v", -1, "amount to send in transaction")

	startTxCmd.AddCommand(startStreamTxCmd, startBurstTxCmd)
	txCmd.AddCommand(sendSingleTxCmd, startTxCmd, stopTxCmd)
	RootCmd.AddCommand(txCmd)
}
