/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add-context CONTEXT_NAME",
	Short: "Add context to config file",
	Long:  `Adds a context describing the eks cluster to the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		clusterName := cmd.Flag("cluster-name").Value.String()
		profile := cmd.Flag("profile").Value.String()
		region := cmd.Flag("region").Value.String()
		currentContext, _ := strconv.ParseBool(cmd.Flag("current-context").Value.String())

		err := NewKubectlEKSConfigImpl().AddContext(name, clusterName, profile, region, currentContext)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	configCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "name of the context")
	addCmd.Flags().StringP("cluster-name", "c", "", "name of EKS cluster")
	addCmd.Flags().StringP("profile", "p", "", "name of AWS profile used to access EKS cluster")
	addCmd.Flags().StringP("region", "r", "", "AWS region cluster is located")
	addCmd.Flags().BoolP("current-context", "s", false, "flag to dictate whether to set newly added context as current context")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("cluster-name")
	addCmd.MarkFlagRequired("profile")
}
