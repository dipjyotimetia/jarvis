package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/dipjyotimetia/jarvis/pkg/engine/gemini"
	"github.com/dipjyotimetia/jarvis/pkg/engine/prompt"
	"github.com/dipjyotimetia/jarvis/pkg/atlassian/confluence"
	"github.com/dipjyotimetia/jarvis/pkg/atlassian/jira"
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

			s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
			s.Color("green")
			s.Suffix = " Generating Tests..."
			s.FinalMSG = "Tests Generated Successfully!\n"

			languageContent := prompt.PromptContent{
				ErrorMsg: "Please provide a valid language.",
				Label:    "What programming lanaguage would you like to use?",
				ItemType: "language",
			}
			language := prompt.SelectLanguage(languageContent)

			specContent := prompt.PromptContent{
				ErrorMsg: "Please provide a valid spec.",
				Label:    "What spec would you like to use?",
				ItemType: "spec",
			}
			spec := prompt.SelectLanguage(specContent)

			file, err := files.ListFiles(specPath)
			if err != nil {
				return fmt.Errorf("failed to identify spec types: %w", err)
			}
			if len(file) == 0 {
				return errors.New("no files found")
			}

			reader, err := files.ReadFile(file[0])
			if err != nil {
				return fmt.Errorf("failed to read spec file: %w", err)
			}

			s.Start()
			ctx := context.Background()
			ai, err := gemini.New(ctx)
			if err != nil {
				return fmt.Errorf("failed to create Gemini engine: %w", err)
			}

			err = ai.GenerateTextStreamWriter(ctx, reader, outputPath, language, spec)
			if err != nil {
				s.FinalMSG = "Test generation failed: %v\n"
				return err
			}
			s.Stop()
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

			specContent := prompt.PromptContent{
				ErrorMsg: "Please provide a valid spec.",
				Label:    "What spec would you like to use?",
				ItemType: "spec",
			}

			spec := prompt.SelectLanguage(specContent)

			ctx := context.Background()
			ai, err := gemini.New(ctx)
			if err != nil {
				return fmt.Errorf("failed to create Gemini engine: %w", err)
			}

			file, err := files.ListFiles(specPath)
			if err != nil {
				return fmt.Errorf("failed to identify spec types: %w", err)
			}
			if len(file) == 0 {
				return errors.New("no files found")
			}

			reader, err := files.ReadFile(file[0])
			if err != nil {
				return fmt.Errorf("failed to read spec file: %w", err)
			}

			err = ai.GenerateTextStream(ctx, reader, spec)
			if err != nil {
				return err
			}
			return nil
		},
	}
	setGenerateScenariosFlag(cmd)
	return cmd
}

func ReadConfluenceJira() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "read-confluence-jira",
		Aliases: []string{"rcj"},
		Short:   "read-confluence-jira is for reading from Confluence and Jira and suggesting test cases.",
		Long:    `read-confluence-jira is for reading from Confluence and Jira and suggesting test cases using Google Gemini`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			// Initialize Confluence client
			confluenceClient := confluence.New(ctx)
			pages, err := confluenceClient.GetPages()
			if err != nil {
				return fmt.Errorf("failed to get pages from Confluence: %w", err)
			}

			// Initialize Jira client
			jiraClient := jira.New(ctx)
			issues, err := jiraClient.GetIssues()
			if err != nil {
				return fmt.Errorf("failed to get issues from Jira: %w", err)
			}

			// Initialize Gemini client
			ai, err := gemini.New(ctx)
			if err != nil {
				return fmt.Errorf("failed to create Gemini engine: %w", err)
			}

			// Combine pages and issues into a single input for Gemini
			input := append(pages, issues...)

			// Generate test cases using Gemini
			err = ai.GenerateTextStream(ctx, input, "confluence-jira")
			if err != nil {
				return fmt.Errorf("failed to generate test cases: %w", err)
			}

			return nil
		},
	}
	return cmd
}
