/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use-context CONTEXT_NAME",
	Short: "Sets the current-context",
	Long:  `Sets the current-context`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}
		NewKubectlEKSConfigImpl().UseContext(args[0])
	},
}

func init() {
	configCmd.AddCommand(useCmd)

}
