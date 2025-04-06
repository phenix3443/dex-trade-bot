/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Server is a command to start the server",
	Long:  `This application is a server to provider services.`,
}
