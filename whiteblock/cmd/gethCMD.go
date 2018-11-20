package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	gethcommand string
)

var gethCmd = &cobra.Command{
	Use:   "geth [command]",
	Short: "Run geth commands",
	Long: `Geth will allow the user to get infromation and run geth commands.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("bash", "-c", "./whiteblock geth -h").Output()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", out)
	},
}

var getBlockNumberCmd = &cobra.Command{
	Use:   "get_block_number",
	Short: "Get block number",
	Long: `Get the current highest block number of the chain
	Response: The block number e.g. 10`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_block_number"
		param := ""
		fmt.Println(serverAddr)
		// fmt.Println(command)
		if len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_block_number -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getBlockCmd = &cobra.Command{
	Use:   "get_block <block number>",
	Short: "Get block information",
	Long: `Get the data of a block
	Response: JSON Representation of the block.
	
	Params: Block number
	Format: <Block Number>`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_block"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_block -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getAccountCmd = &cobra.Command{
	Use:   "get_accounts",
	Short: "Get account information",
	Long: `Get a list of all unlocked accounts
	Response: A JSON array of the accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_accounts"
		param := ""
		// fmt.Println(command)
		if len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_accounts -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getBalanceCmd = &cobra.Command{
	Use:   "get_balance <address>",
	Short: "Get account balance information",
	Long: `Get the current balance of an account
	Response: The integer balance of the account in wei
	
	Params: Account address
	Format: <address>`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_balance"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_balance -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var sendTxCmd = &cobra.Command{
	Use:   "send_transaction <from address> <to address> <gas> <gas price> <value to send>",
	Short: "Sends a transaction",
	Long: `Send a transaction between two accounts
	Response: The transaction hash

	Params: Sending account, receiving account, gas, gas price, amount to send, transaction data, nonce
	Format: <from> <to> <gas> <gas price> <value>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::send_transaction"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) <= 4 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth send_transaction -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getTxCountCmd = &cobra.Command{
	Use:   "get_transaction_count <address> [block number]",
	Short: "Get transaction count",
	Long: `Get the transaction count sent from an address, optionally by block
	Response: The transaction count
	
	Params: The sender account, a block number
	Format: <address> [block number]`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_transaction_count"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 2 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_transaction_count -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getTxCmd = &cobra.Command{
	Use:   "get_transaction <hash>",
	Short: "Get transaction information",
	Long: `Get a transaction by its hash
	Response: JSON representation of the transaction.
	
	Params: The transaction hash
	Format: <hash>`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_transaction"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_transaction -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getTxReceiptCmd = &cobra.Command{
	Use:   "get_transaction_receipt <hash>",
	Short: "Get transaction receipt",
	Long: `Get the transaction receipt by the tx hash
	Response: JSON representation of the transaction receipt.
	
	Params: The transaction hash
	Format: <hash>`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_transaction_receipt"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_transaction_receipt -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getHashRateCmd = &cobra.Command{
	Use:   "get_hash_rate",
	Short: "Get hasg rate",
	Long: `Get the current hash rate per node
	Response: The hash rate of a single node in the network`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_hash_rate"
		param := ""
		// fmt.Println(command)
		if len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_hash_rate -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var startTxCmd = &cobra.Command{
	Use:   "start_transactions <tx/s> <value> [destination]",
	Short: "Start transactions",
	Long: `Start sending transactions according to the given parameters, value = -1 means randomize value.
	
	Params: The amount of transactions to send in a second, the value of each transaction in wei, the destination for the transaction
	Format: <tx/s> <value> [destination]`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::start_transactions"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) <= 1 || len(args) > 3 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth start_transactions -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var stopTxCmd = &cobra.Command{
	Use:   "stop_transactions",
	Short: "Stop transactions",
	Long:  `Stops the sending of transactions if transactions are currently being sent`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::stop_transactions"
		param := ""
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth stop_transactions -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var startMiningCmd = &cobra.Command{
	Use:   "start_mining [node 1 number] [node 2 number]...",
	Short: "Start Mining",
	Long: `Send the start mining signal to nodes, may take a while to take effect due to DAG generation
	Response: The number of nodes which successfully received the signal to start mining
	
	Params: A list of the nodes to start mining or None for all nodes
	Format: [node 1 number] [node 2 number]...`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::start_mining"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth start_mining -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var stopMiningCmd = &cobra.Command{
	Use:   "stop_mining [node 1 number] [node 2 number]...",
	Short: "Stop mining",
	Long: `Send the stop mining signal to nodes
	Response: The number of nodes which successfully received the signal to stop mining
	
	Params: A list of the nodes to stop mining or None for all nodes
	Format: [node 1 number] [node 2 number]...`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::stop_mining"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth stop_mining -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var blockListenerCmd = &cobra.Command{
	Use:   "block_listener [block number]",
	Short: "Get block listener",
	Long: `Get all blocks and continue to subscribe to new blocks
	Response: Will emit on eth::block_listener for every block after the given block or 0 that exists/has been created
	
	Params: The block number to start at or None for all blocks
	Format: [block number]`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::block_listener"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth block_listener -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

var getRecentSentTxCmd = &cobra.Command{
	Use:   "get_recent_sent_tx [number]",
	Short: "Get recently sent transaction",
	Long: `Get a number of the most recent transactions sent
	
	Response: JSON object of transaction data
	
	Params: The number of transactions to retrieve
	Format: [number]`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "eth::get_recent_sent_tx"
		param := strings.Join(args[:], " ")
		// fmt.Println(command)
		if len(args) == 0 || len(args) > 1 {
			out, err := exec.Command("bash", "-c", "./whiteblock geth get_recent_sent_tx -h").Output()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", out)
		}
		wsGethCmd(serverAddr, command, param)
	},
}

func init() {
	// gethCmd.Flags().StringVarP(&gethcommand, "command", "c", "", "Geth command")
	gethCmd.Flags().StringVarP(&serverAddr, "server-addr", "a", "localhost:5000", "server address with port 5000")

	//geth subcommands

	gethCmd.AddCommand(getBlockNumberCmd)
	gethCmd.AddCommand(getBlockCmd)
	gethCmd.AddCommand(getAccountCmd)
	gethCmd.AddCommand(getBalanceCmd)
	gethCmd.AddCommand(sendTxCmd)
	gethCmd.AddCommand(getTxCountCmd)
	gethCmd.AddCommand(getTxCmd)
	gethCmd.AddCommand(getTxReceiptCmd)
	gethCmd.AddCommand(getHashRateCmd)
	gethCmd.AddCommand(startTxCmd)
	gethCmd.AddCommand(stopTxCmd)
	gethCmd.AddCommand(startMiningCmd)
	gethCmd.AddCommand(stopMiningCmd)
	gethCmd.AddCommand(blockListenerCmd)
	gethCmd.AddCommand(getRecentSentTxCmd)

	RootCmd.AddCommand(gethCmd)
}
