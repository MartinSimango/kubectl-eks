/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the current config",
	Long:  `Displays the config file used for the eks kubectl plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		displayConfig()
	},
}

func init() {
	configCmd.AddCommand(viewCmd)
}

func displayConfig() {
	config := NewKubectlEKSConfigImpl().GetConfig()

	yamlData, err := yaml.Marshal(config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(yamlData[:]))

}
