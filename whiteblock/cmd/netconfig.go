package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var netropyCmd = &cobra.Command{
	Use:     "netconfig <engine number> <path number> <engine number> <path number> <command>",
	Aliases: []string{"emulate"},
	Short:   "Network conditions",
	Long: `
Netconfig will introduce persisting network conditions for testing.
	
	latency 			Specifies the latency to add [ms];
	packetloss 			Specifies the amount of packet loss to add [%];
	bandwidth			Specifies the bandwidth of the network [bps|Kbps|Mbps|Gbps];
	`,

	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "netconfig"
		msg := "engine " + args[0] + " path " + args[1] + " " + strings.Join(args[4:], " ")

		wsSendCmd(serverAddr, command, msg)
	},
}

var latencyCmd = &cobra.Command{
	Use:     "latency <engine number> <path number> <engine number> <path number>",
	Aliases: []string{"lat"},
	Short:   "Set latency",
	Long: `
Latency will introduce delay to the network.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "netconfig"
		msg1 := "engine " + args[0] + " path " + args[1] + " set delay constant port 1 to port 2"
		msg2 := "engine " + args[2] + " path " + args[3] + " set delay constant port 2 to port 1"

		wsSendCmd(serverAddr, command, msg1)
		wsSendCmd(serverAddr, command, msg2)
	},
}

var packetLossCmd = &cobra.Command{
	Use:     "packetloss <engine number> <path number> <engine number> <path number> <percent>",
	Aliases: []string{"loss"},
	Short:   "Set packetloss",
	Long: `
Packetloss will drop packets in the network.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "netconfig"
		msg1 := "engine " + args[0] + " path " + args[1] + " set loss random " + args[4] + " port 1 to port 2"
		msg2 := "engine " + args[2] + " path " + args[3] + " set loss random " + args[4] + " port 2 to port 1"

		wsSendCmd(serverAddr, command, msg1)
		wsSendCmd(serverAddr, command, msg2)
	},
}

var bandwCmd = &cobra.Command{
	Use:     "bandwidth <engine number> <path number> <engine number> <path number> <bandwidth number> <bandwidth type>",
	Aliases: []string{"bw"},
	Short:   "Set bandwidth",
	Long: `
Bandwidth will constrict the network to the specified bandwidth.

Fomat: 
	bandwidth type: bps, Kbps, Mbps, Gbps
	`,

	Run: func(cmd *cobra.Command, args []string) {
		serverAddr = "ws://" + serverAddr + "/socket.io/?EIO=3&transport=websocket"
		command := "netconfig"
		msg1 := "engine " + args[0] + " path " + args[1] + " set bw fixed " + args[4] + args[5]
		msg2 := "engine " + args[2] + " path " + args[3] + " set bw fixed " + args[4] + args[5]

		wsSendCmd(serverAddr, command, msg1)
		wsSendCmd(serverAddr, command, msg2)
	},
}

func init() {
	netropyCmd.AddCommand(latencyCmd, packetLossCmd, bandwCmd)

	RootCmd.AddCommand(netropyCmd)
}
