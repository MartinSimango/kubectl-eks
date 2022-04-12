/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// resyncCmd represents the reload command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs kubectl with the current context",
	Long:  `Syncs kubectl with the current context`,
	Run: func(cmd *cobra.Command, args []string) {
		NewKubectlEKSConfigImpl().Sync()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
