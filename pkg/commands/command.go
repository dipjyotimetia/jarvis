package commands

import (
	"context"

	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/dipjyotimetia/jarvis/pkg/engine/gemini"
	"github.com/spf13/cobra"
)

func setGenerateTestFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "", "spec path")
	cmd.Flags().StringP("output", "o", "", "output path")
}

func setGenerateScenariosFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "", "spec path")
}

func GenerateTestModule() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-test",
		Aliases: []string{"test"},
		Short:   "generate-test is for generating test cases.",
		Long:    `generate-test is for generating test cases from the provided spec files`,
		RunE: func(cmd *cobra.Command, args []string) error {
			specPath, _ := cmd.Flags().GetString("path")
			outputPath, _ := cmd.Flags().GetString("output")

			ctx := context.Background()

			ai, err := gemini.NewGenClient(ctx)
			if err != nil {
				panic(err)
			}

			file, err := files.IdentifySpecTypes(specPath)
			if err != nil {
				panic(err)
			}
			reader, err := files.ReadFile(file[0])
			if err != nil {
				panic(err)
			}
			err = ai.GenerateTextStreamWriter(ctx, reader, outputPath)
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
	setGenerateTestFlag(cmd)
	return cmd
}

func GenerateTestScenarios() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-scenarios",
		Aliases: []string{"scenarios"},
		Short:   "generate-scenarios is for generating test scenarios.",
		Long:    `generate-scenarios is for generating test scenarios from the provided spec files`,
		RunE: func(cmd *cobra.Command, args []string) error {
			specPath, _ := cmd.Flags().GetString("path")

			ctx := context.Background()

			ai, err := gemini.NewGenClient(ctx)
			if err != nil {
				panic(err)
			}

			file, err := files.IdentifySpecTypes(specPath)
			if err != nil {
				panic(err)
			}
			reader, err := files.ReadFile(file[0])
			if err != nil {
				panic(err)
			}
			err = ai.GenerateTextStream(ctx, reader)
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
	setGenerateScenariosFlag(cmd)
	return cmd
}
