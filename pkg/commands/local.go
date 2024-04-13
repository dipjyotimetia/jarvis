package commands

import (
	"errors"
	"fmt"

	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/dipjyotimetia/jarvis/pkg/engine/prompt"
	"github.com/dipjyotimetia/jarvis/pkg/engine/utils"
	"github.com/spf13/cobra"
)

func setSpecPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "", "spec path")
}

func setGrpCurlPathFlag(cmd *cobra.Command) {
	cmd.Flags().String("proto", "", "protofile path")
	cmd.Flags().String("service", "", "service name")
	cmd.Flags().String("method", "", "method name")
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
			fmt.Println("Analyzing spec files..." + spec)
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

func GrpcCurlGenerator() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "grpc-curl",
		Aliases: []string{"grpc-curl"},
		Short:   "grpc-curl is for generating curl commands for grpc services",
		Long:    `grpc-curl is for generating curl commands for grpc services`,
		RunE: func(cmd *cobra.Command, args []string) error {
			protoFile, _ := cmd.Flags().GetString("proto")
			serviceName, _ := cmd.Flags().GetString("service")
			methodName, _ := cmd.Flags().GetString("method")
			utils.GrpCurlCommand(protoFile, serviceName, methodName)
			return nil
		},
	}
	setGrpCurlPathFlag(cmd)
	return cmd
}
