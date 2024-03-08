package commands

import (
	"errors"

	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/dipjyotimetia/jarvis/pkg/engine/prompt"
	"github.com/dipjyotimetia/jarvis/pkg/engine/utils"
	"github.com/spf13/cobra"
)

func setSpecPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "", "spec path")
}

func SpecAnalyzer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "spec-analyzer",
		Aliases: []string{"spec"},
		Short:   "proto-analyzer is for analyzing protobuf spec files",
		Long:    `proto-analyzer is for analyzing protobuf spec files`,
		RunE: func(cmd *cobra.Command, args []string) error {
			specPath, _ := cmd.Flags().GetString("path")

			specContent := prompt.PromptContent{
				ErrorMsg: "Please provide a valid spec.",
				Label:    "What spec would you like to use?",
				ItemType: "spec",
			}
			spec := prompt.SelectLanguage(specContent)

			specs, err := files.ListFiles(specPath)
			if err != nil {
				return err
			}
			if len(specs) == 0 {
				return errors.New("no files found")
			}
			switch spec {
			case "protobuf":
				utils.ProtoAnalyzer(specs)
			case "openapi":
				utils.OpenApiAnalyzer(specs)
			default:
				return nil
			}
			return nil
		},
	}
	setSpecPathFlag(cmd)
	return cmd
}
