package cmd

import (
	"fmt"
	"os"

	"github.com/dipjyotimetia/jarvis/pkg/commands"
	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:     "jarvis",
	Short:   "A generative AI-driven CLI for testing",
	Long:    `Jarvis is a powerful toolkit that leverages generative AI to streamline and enhance various testing activities.`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Dipjyoti Metia", "")
	rootCmd.AddCommand(commands.GenerateTestModule())
	rootCmd.AddCommand(commands.GenerateTestScenarios())
	rootCmd.AddCommand(commands.SpecAnalyzer())
	rootCmd.AddCommand(commands.GrpcCurlGenerator())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
