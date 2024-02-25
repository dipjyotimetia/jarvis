package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
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
		fmt.Println(d.Name())
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

func ReadFile(spec string) ([]genai.Text, error) {
	file, err := os.Open(spec)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()
	var lines []genai.Text

	reader := bufio.NewReader(file)
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)

	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Errorf("Error reading file:", err)
			}
			break
		}
		lines = append(lines, genai.Text(buffer[:bytesRead]))
	}
	return lines, nil
}

func CheckDirectryExists(output string) {
	if _, err := os.Stat(fmt.Sprintf("./%s", output)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("./%s", output), 0755)
	}
}
