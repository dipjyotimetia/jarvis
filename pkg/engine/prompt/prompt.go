package prompt

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"github.com/manifoldco/promptui"
)

type specType string

const AVRO specType = "avro"
const PROTOBUF specType = "protobuf"
const SWAGGER specType = "swagger"
const OAS3 specType = "OpenAPI 3"

type PromptContent struct {
	ErrorMsg string
	Label    string
	ItemType string
	Items    []string
}

// CompareSpecFiles returns a prompt to compare two spec files
func CompareSpecFiles(file1, file2 genai.Part, specType specType) []genai.Part {
	prompt := []genai.Part{
		file1,
		file2,
		genai.Text(fmt.Sprintf("compare the two %s contracts, check if there a breaking change introduced by this change?", specType)),
	}

	return prompt
}

func SelectLanguage(pc PromptContent) string {
	var items []string
	switch pc.ItemType {
	case "language":
		items = []string{"Go", "Python", "JavaScript", "java", "TypeScript"}
	case "spec":
		items = []string{"avro", "protobuf", "swagger", "openapi"}
	default:
		items = []string{}
		fmt.Println("Invalid prompt type")
	}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		color.Red(pc.ErrorMsg)
		os.Exit(1)
	}

	color.Green("Input: %s\n", result)

	return result
}

func setFrameworksForLanguage(language string) []string {
	switch language {
	case "Go":
		return []string{"Gin", "Echo", "Fiber", "gRPC"}
	case "Python":
		return []string{"Django", "Flask", "FastAPI", "gRPC"}
	case "Java":
		return []string{"Spring", "JAX-RS", "restassured", "gRPC"}
	case "JavaScript":
		return []string{"supertest", "axios", "http", "gRPC"}
	case "TypeScript":
		return []string{"supertest", "axios", "http"}
	default:
		return []string{}
	}
}
