/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete-context",
	Short: "Delete context from config file",
	Long:  `Deletes a context describing the eks cluster from the config file`,
	Run: func(command *cobra.Command, args []string) {
		NewKubectlEKSConfigImpl().DeleteContext(args[0])
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)

}
