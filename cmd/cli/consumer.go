/*
Copyright © 2025 Beingana Jim Junior
*/
package cli

import (
	"fmt"

	csmr "github.com/jim-junior/eda/cmd/consumer"
	"github.com/spf13/cobra"
)

var dfe bool

var isMaster bool

var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the Consumer",
	Run: func(cmd *cobra.Command, args []string) {

		if dfe {

			if isMaster {
				csmr.RunMasterConsumer()
			} else {
				csmr.RunDFEConsumer()
			}

		} else {
			topic := cmd.Flag("topic").Value.String()
			if topic == "" {
				fmt.Println("ERROR: You must specify the Topic to listen too. use --topic or -t flags to specify it")
				return
			}
			csmr.RunConsumer(topic)
		}
	},
}

func init() {
	ConsumerCmd.Flags().StringP("topic", "t", "", "Topic for the consumer to listen too")
	//ConsumerCmd.Flags().StringP("deffered-exec", "d", "", "Run Consumer to demostrate Deferred Execution and Eventual Consistency")

	ConsumerCmd.Flags().BoolVarP(&dfe, "deffered-exec", "d", false, "Run Consumer to demostrate Deferred Execution and Eventual Consistency")

	ConsumerCmd.Flags().BoolVarP(&isMaster, "master", "m", false, "Run Master Node")
}
