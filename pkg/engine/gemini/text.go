package gemini

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

// GenerateText GenerateContent
func (c *client) GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	resp, err := c.ProModel().GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GenerateTextStream GenerateContentStream
func (c *client) GenerateTextStream(ctx context.Context, specs []genai.Text, specType string) error {
	var prompts []genai.Part
	for _, spec := range specs {
		prompts = append(prompts, spec)
	}
	prompts = append(prompts, genai.Text(fmt.Sprintf("Generate all possible positive and negative test scenarios in simple english for the provided %s spec file.", specType)))

	resp := c.ProModel().GenerateContentStream(ctx, prompts...)
	for {
		resp, err := resp.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}
		if resp.Candidates == nil {
			return nil
		}
		for _, candidate := range resp.Candidates {
			go func(parts []genai.Part) {
				for _, c := range parts {
					fmt.Println(c)
				}
			}(candidate.Content.Parts)
		}
	}
	return nil
}

// GenerateTextStreamWriter GenerateContentStream
func (c *client) GenerateTextStreamWriter(ctx context.Context, specs []genai.Text, language, specType string, outputFolder string) error {
	var prompts []genai.Part
	for _, spec := range specs {
		prompts = append(prompts, spec)
	}
	prompts = append(prompts, genai.Text(fmt.Sprintf("Generate %s tests based on this %s spec.", language, specType)))

	ct := time.Now().Format("2006-01-02-15-04-05")
	files.CheckDirectryExists(outputFolder)
	outputFile, err := os.Create(fmt.Sprintf("%s/%s_output_test.md", outputFolder, ct))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	response := c.ProModel().GenerateContentStream(ctx, prompts...)
	for {
		resp, err := response.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}
		if resp.Candidates == nil {
			return nil
		}
		for _, candidate := range resp.Candidates {
			for _, c := range candidate.Content.Parts {
				_, err := fmt.Fprintln(writer, c)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
