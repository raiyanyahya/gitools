package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gts",
	Short: "A cli to help you when you are stuck.",
	Long:  `A command line application to help when you are stuck in tricky git related scenarios.`,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
