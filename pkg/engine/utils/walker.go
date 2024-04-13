package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/olekukonko/tablewriter"
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

func ProtoAnalyzer(protoFiles []string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File", "Service", "Method", "Input Type", "Output Type", "Streaming"})
	for _, protoFile := range protoFiles {
		parser := protoparse.Parser{
			//required for google proto files
			ImportPaths:           []string{"."},
			IncludeSourceCodeInfo: true,
			InferImportPaths:      true,
		}
		fds, err := parser.ParseFiles(protoFile)
		if err != nil {
			return fmt.Errorf("error parsing Proto file %s: %v", protoFile, err)
		}

		for _, file := range fds {
			for _, service := range file.GetServices() {
				for _, method := range service.GetMethods() {
					descriptor := method.AsMethodDescriptorProto()
					streaming := "No"
					if descriptor.GetClientStreaming() || descriptor.GetServerStreaming() {
						streaming = "Yes"
					}

					table.Append([]string{
						file.GetName(),
						service.GetName(),
						method.GetName(),
						descriptor.GetInputType(),
						descriptor.GetOutputType(),
						streaming,
					})
				}
			}
		}
	}

	table.Render()
	return nil
}

// generateGrpcurlCommand generates a grpcurl command for a given service and method
func GrpCurlCommand(protoFile, serviceName, methodName string) {
	var grpCurl string
	parser := protoparse.Parser{
		//required for google proto files
		ImportPaths:           []string{"."},
		IncludeSourceCodeInfo: true,
		InferImportPaths:      true,
	}

	fds, err := parser.ParseFiles(protoFile)
	if err != nil {
		fmt.Errorf("error parsing Proto file %s: %v", protoFile, err)
	}
	for _, file := range fds {
		for _, service := range file.GetServices() {
			if service.GetName() == serviceName {
				for _, method := range service.GetMethods() {
					if method.GetName() == methodName {
						grpCurl = fmt.Sprintf("grpcurl -plaintext -proto %s -d '{\"%s\"}' localhost:50051 %s/%s",
							"", method.AsMethodDescriptorProto().GetInputType(), service.GetFullyQualifiedName(), methodName)
					}
				}
			}
		}
	}
	fmt.Println(grpCurl)
}
