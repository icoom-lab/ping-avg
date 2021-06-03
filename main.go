package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

func main() {
	var count int

	var rootCmd = &cobra.Command{
		Use:   "ping-avg",
		Short: "Ping avg",
		Long:  `Ping multiple hosts to determine the average over time`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			w := tabwriter.NewWriter(os.Stdout, 10, 10, 3, ' ', 0)
			defer w.Flush()

			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "ip", "avg", "min", "max", "loss")
			fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")

			for _, ip := range args {

				pinger, err := ping.NewPinger(ip)
				if err != nil {
					panic(err)
				}

				pinger.Count = count

				err = pinger.Run()
				if err != nil {
					panic(err)
				}
				stats := pinger.Statistics()
				fmt.Fprintf(w, "\n %v\t%v\t%v\t%v\t%v\t", ip, stats.AvgRtt, stats.MinRtt, stats.MaxRtt, stats.PacketLoss)
			}
		},
	}

	rootCmd.Flags().IntVarP(&count, "count", "c", 4, "Specifies the number of echo Request messages be sent. The default is 4.")

	rootCmd.Execute()
}
