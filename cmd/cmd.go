package cmd

import (
	"fmt"
	"os"

	"github.com/dipjyotimetia/jarvis/pkg/commands"
	"github.com/spf13/cobra"
)

var version = "0.0.2"

var rootCmd = &cobra.Command{
	Use:     "jarvis",
	Short:   "jarvis is a very fast static site generator",
	Version: version,
	Long:    `Generative ai based cli to perform different testing activity`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Dipjyoti Metia", "")
	rootCmd.AddCommand(commands.GenerateTestModule())
	rootCmd.AddCommand(commands.GenerateTestScenarios())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
