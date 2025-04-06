/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/phenix3443/dex-trade-bot/pkg/api/router"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Gateway is a command to start the gateway server",
	RunE: func(_ *cobra.Command, _ []string) error {
		r := router.SetupRouter()
		return r.Run(appConfig.Gateway.Address)
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}
