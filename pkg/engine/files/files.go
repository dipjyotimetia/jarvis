package files

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func ListFiles(dir string) ([]string, error) {
	paths := []string{}
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		paths = append(paths, path)

		// Read the file
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return nil
		}

		// Identify the file type
		fileType := identifyFileType(data)
		fmt.Printf("File: %s, Type: %s\n", d.Name(), fileType)

		return nil
	})
	return paths, nil
}

func identifyFileType(data []byte) string {
	// OpenAPI (JSON) Check
	var openAPI interface{}
	if err := json.Unmarshal(data, &openAPI); err == nil {
		if _, ok := openAPI.(map[string]interface{})["openapi"]; ok {
			return "OpenAPI (JSON)"
		}
	}

	// OpenAPI (YAML) Check
	if strings.Contains(string(data), "openapi:") || strings.Contains(string(data), "swagger:") {
		return "OpenAPI (YAML)"
	}

	// Protobuf Check (Very basic - Needs hints about expected message types)
	if regexp.MustCompile(`(?m)^(package|message|syntax)\s`).Match(data) {
		return "Protobuf (Likely)"
	}

	// Avro Check (Rudimentary)
	if strings.Contains(string(data), "{\"type\": \"record\"") {
		return "Avro (Likely)"
	}

	return "Unknown"
}

func ReadFile(path string) ([]genai.Text, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var texts []genai.Text
	for _, line := range lines {
		texts = append(texts, genai.Text(line))
	}

	return texts, nil
}

func CheckDirectryExists(output string) {
	if _, err := os.Stat(fmt.Sprintf("./%s", output)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("./%s", output), 0755)
	}
}
