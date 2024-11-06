package main

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{}

var upgradeNodePackageCmd = &cobra.Command{
	Use:   "upgrade-node-package",
	Short: "Upgrade node package",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upgradeNodePackageCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
