package cmd

import (
	"github.com/spf13/cobra"
)

var eosCmd = &cobra.Command{
	Use:   "eos <command>",
	Short: "Run eos commands",
	Long: `
Eos will allow the user to get infromation and run geth commands.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		return
	},
}

var eosGetBlockCmd = &cobra.Command{
	Use:   "get_block <block number>",
	Short: "Get block number",
	Long: `
Roughly equivalent to calling cleos get block <block number>

Params: The block number
Format: <block number>

Response: Block data for that block`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::get_block"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) != 1 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

var eosGetInfoCmd = &cobra.Command{
	Use:   "get_info [node]",
	Short: "Get EOS info",
	Long: `
Roughly equivalent to calling cleos get info

Params: The node to get info from
Format: [node]

Response: EOS Info`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::get_info"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) > 1 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

var eosSendTxCmd = &cobra.Command{
	Use:   "send_transaction <node> <from> <to> <amount> [symbol=SYS] [code=eosio.token] [memo=]",
	Short: "Send single transaction to another account",
	Long: `
This command will send a single transaction from one account to another. Additional arguments are required.

Params: node number, from account, to account, amount, symbol, code, memo
Format: <node> <from> <to> <amount> [symbol=SYS] [code=eosio.token] [memo=]

Response: EOS Info`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::send_transaction"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) < 4 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

var eosSendBurstTxCmd = &cobra.Command{
	Use:   "send_burst_transaction <node> <from> <to> <amount> [symbol=SYS] [code=eosio.token] [memo=]",
	Short: "Send burst transactions",
	Long: `
This command will send a burst of transactions from one account to another. Additional arguments are required.

Params: node number, from address, to address, amount, symbol, code, memo
Format: <node> <from> <to> <amount> [symbol=SYS] [code=eosio.token] [memo=]

Response: EOS Info`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::run_burst_tx"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) < 4 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

var eosConstTpsCmd = &cobra.Command{
	Use:   "run_constant_tps <tps>",
	Short: "Send continuous transactions",
	Long: `
This command will have all nodes send continous transactions at a constant rate..

Params: number to transactions per second
Format: <tps>

Response: success or error`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::run_constant_tps"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) != 1 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

var eosGetBlockNumCmd = &cobra.Command{
	Use:   "get_block_number [node]",
	Short: "Send continuous transactions",
	Long: `
This command will have all nodes send continous transactions at a constant rate..

Params: The node to get it from, default is 0
Format: [node]

Response: Data on the last x test results`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eos::get_block_number"
		param := args[0]
		// fmt.Println(command, param)
		if len(args) > 1 {
			println("\nError: Invalid number of arguments given\n")
			cmd.Help()
			return
		}
		wsEmitListen(serverAddr, command, param)
	},
}

func init() {
	// gethCmd.Flags().StringVarP(&gethcommand, "command", "c", "", "Geth command")
	eosCmd.Flags().StringVarP(&serverAddr, "server-addr", "a", "localhost:5000", "server address with port 5000")

	//geth subcommands
	eosCmd.AddCommand(eosGetBlockCmd, eosGetInfoCmd, eosSendTxCmd, eosSendBurstTxCmd, eosConstTpsCmd, eosGetBlockNumCmd)

	RootCmd.AddCommand(eosCmd)
}
