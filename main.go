package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
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
				count, _ := cmd.Flags().GetInt("count")
				verbose, _ := cmd.Flags().GetBool("verbose")

				pinger.Count = count
				pinger.Timeout = time.Duration(1.15*float64(count+1)) * time.Second

				if verbose {
					pinger.OnRecv = func(pkt *ping.Packet) {
						fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
							pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
					}

					pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
						fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
							pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
					}
				}

				err = pinger.Run()
				if err != nil {
					panic(err)
				}
				stats := pinger.Statistics()
				fmt.Fprintf(w, "\n %v\t%v\t%v\t%v\t%v\t", ip, stats.AvgRtt, stats.MinRtt, stats.MaxRtt, stats.PacketLoss)
			}
		},
	}

	rootCmd.Flags().IntP("count", "c", 1, "Specifies the number of echo Request messages be sent. The default is 1.")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")

	rootCmd.Execute()
}
