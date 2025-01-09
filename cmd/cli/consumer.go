/*
Copyright © 2025 Beingana Jim Junior
*/
package cli

import (
	csmr "github.com/jim-junior/eda/cmd/consumer"
	"github.com/spf13/cobra"
)

var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		csmr.RunConsumer()
	},
}
