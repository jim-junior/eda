/*
Copyright © 2025 Beingana Jim Junior
*/
package cli

import (
	apiServer "github.com/jim-junior/eda/cmd/api"
	"github.com/spf13/cobra"
)

var APIServerCmd = &cobra.Command{
	Use:   "api-server",
	Short: "Start the API Server",
	Run: func(cmd *cobra.Command, args []string) {
		apiServer.StartServer()
	},
}
