package prompt

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

type specType string

const AVRO specType = "avro"
const PROTOBUF specType = "protobuf"
const SWAGGER specType = "swagger"
const OAS3 specType = "OpenAPI 3"

// CompareSpecFiles returns a prompt to compare two spec files
func CompareSpecFiles(file1, file2 genai.Part, specType specType) []genai.Part {
	prompt := []genai.Part{
		file1,
		file2,
		genai.Text(fmt.Sprintf("compare the two %s contracts, check if there a breaking change introduced by this change?", specType)),
	}

	return prompt
}
