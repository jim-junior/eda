/*
Copyright © 2024 Cranom Technologies Limited info@cranom.tech
*/
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eda",
	Short: "CLI Tool to demostrate Event Driven Architecture.",
	Long: `CLI Tool to demostrate Event Driven Architecture.
	
	Learn More at: https://jim-junior.github.io/#/blog/a-practical-guide-to-the-event-driven-architecture
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `eda api-server` for to start the api server and then run `eda consumer` to start the consumer")
		fmt.Println("Use eda --help to see available commands")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(APIServerCmd)
	rootCmd.AddCommand(ConsumerCmd)

}
