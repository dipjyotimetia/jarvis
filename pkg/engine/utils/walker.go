package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"gopkg.in/yaml.v3"
)

type OpenAPI struct {
	Paths map[string]PathItem `json:"paths"`
}

type Operation struct {
	OperationID string `json:"operationId"`
}

type PathItem struct {
	Get    *Operation `json:"get"`
	Post   *Operation `json:"post"`
	Put    *Operation `json:"put"`
	Delete *Operation `json:"delete"`
	Patch  *Operation `json:"patch"`
}

func OpenApiAnalyzer(specFiles []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Method", "Path", "OperationID"})

	for _, specFile := range specFiles {
		data, err := os.ReadFile(specFile)
		if err != nil {
			panic(err)
		}

		var openapi OpenAPI
		if err := json.Unmarshal(data, &openapi); err == nil {
		} else if err := yaml.Unmarshal(data, &openapi); err == nil {
		} else {
			panic("Unsupported OpenAPI file format")
		}

		for path, pathItem := range openapi.Paths {
			if pathItem.Get != nil {
				table.Append([]string{"GET", path, pathItem.Get.OperationID})
			}
			if pathItem.Post != nil {
				table.Append([]string{"POST", path, pathItem.Post.OperationID})
			}
			if pathItem.Put != nil {
				table.Append([]string{"PUT", path, pathItem.Put.OperationID})
			}
			if pathItem.Patch != nil {
				table.Append([]string{"PATCH", path, pathItem.Patch.OperationID})
			}
			if pathItem.Delete != nil {
				table.Append([]string{"DELETE", path, pathItem.Delete.OperationID})
			}
		}
	}
	table.Render()
}

func ProtoAnalyzer(protoFiles []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File", "Service", "Method", "Input Type", "Output Type", "Streaming"})

	for _, protoFile := range protoFiles {
		data, err := os.ReadFile(protoFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		fds := &descriptorpb.FileDescriptorSet{}
		if err := proto.Unmarshal(data, fds); err != nil {
			fmt.Println("Error parsing Proto:", err)
			continue
		}

		for _, file := range fds.File {
			for _, service := range file.Service {
				for _, method := range service.Method {
					streaming := "No"
					if method.GetClientStreaming() || method.GetServerStreaming() {
						streaming = "Yes"
					}

					table.Append([]string{
						file.GetName(),
						service.GetName(),
						method.GetName(),
						method.GetInputType(),
						method.GetOutputType(),
						streaming,
					})
				}
			}
		}
	}

	table.Render()
}
