package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

func main() {
	var count int

	var rootCmd = &cobra.Command{
		Use:   "ping-avg",
		Short: "Ping avg",
		Long:  `Ping multiple hosts to determine the average over time`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, ip := range args {

				pinger, err := ping.NewPinger(ip)
				if err != nil {
					panic(err)
				}

				pinger.Count = count

				err = pinger.Run() // Blocks until finished.
				if err != nil {
					panic(err)
				}
				stats := pinger.Statistics()

				fmt.Printf("%v 	avg=%v 		min=%v 	max=%v 	loss=%v\n", ip, stats.AvgRtt, stats.MinRtt, stats.MaxRtt, stats.PacketLoss)
			}
		},
	}

	rootCmd.Flags().IntVarP(&count, "count", "c", 4, "Specifies the number of echo Request messages be sent. The default is 4.")

	rootCmd.Execute()
}
