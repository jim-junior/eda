/*
Copyright © 2025 Beingana Jim Junior
*/
package cli

import (
	"fmt"

	csmr "github.com/jim-junior/eda/cmd/consumer"
	"github.com/spf13/cobra"
)

var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		topic := cmd.Flag("topic").Value.String()
		if topic == "" {
			fmt.Println("ERROR: You must specify the Topic to listen too. use --topic or -t flags to specify it")
			return
		}
		csmr.RunConsumer(topic)
	},
}

func init() {
	ConsumerCmd.Flags().StringP("topic", "t", "", "Topic for the consumer to listen too")
}
