/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// currentContextCmd represents the currentContext command
var currentContextCmd = &cobra.Command{
	Use:   "current-context CONTEXT_NAME",
	Short: "Gets the current context ",
	Long:  `Displays the current context the kubectl-eks plugin is using`,
	Run: func(cmd *cobra.Command, args []string) {
		displayCurrentContext()
	},
}

func init() {
	configCmd.AddCommand(currentContextCmd)
}

func displayCurrentContext() {
	fmt.Println("current-context: " + NewKubectlEKSConfigImpl().CurrentContext)
}
