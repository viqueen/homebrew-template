package main

import (
	"github.com/spf13/cobra"
	upgradenodepackage "homebrew/internal/upgrade-node-package"
	"log"
)

var rootCmd = &cobra.Command{}

var upgradeNodePackageCmd = &cobra.Command{
	Use:   "upgrade-node-package",
	Short: "Upgrade node package",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return upgradenodepackage.Task(upgradenodepackage.PackageInfo{
			Org:  args[0],
			Name: args[1],
		})
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
