/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify kubectl-eks config files using subcommands like 'kubectl eks config delete-context my-context' ",
	Long:  `Modify kubectl-eks config files using subcommands like 'kubectl eks config delete-context my-context'`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
