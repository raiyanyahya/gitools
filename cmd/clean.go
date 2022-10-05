package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up and start fresh from master",
	Long:  "This cleans up EVERYTHING and gets your local repo in sync with master.",
	Run:   clean,
}

func clean(cmd *cobra.Command, args []string) {
	fmt.Println("removing changes and moving to a fresh copy of the master branch...")
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
