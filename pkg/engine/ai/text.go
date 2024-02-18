package ai

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func (c *GenClient) GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	resp, err := c.ProModel().GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GenerateContentStream
func (c *GenClient) GenerateTextStream(ctx context.Context, specs []genai.Text) (*genai.GenerateContentResponseIterator, error) {
	prompts := []genai.Part{}
	for _, spec := range specs {
		prompts = append(prompts, spec)
	}
	prompts = append(prompts, genai.Text("Create golang grpc integration tests based on the above proto file."))

	resp := c.ProModel().GenerateContentStream(ctx, prompts...)
	for {
		resp, err := resp.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if resp.Candidates == nil {
			return nil, nil
		}
		for _, candidate := range resp.Candidates {
			for _, c := range candidate.Content.Parts {
				fmt.Println(c)
			}
		}
	}
	return nil, nil
}

func (c *GenClient) GenerateTextStreamWriter(ctx context.Context, specs []genai.Text) error {
	prompts := []genai.Part{}
	for _, spec := range specs {
		prompts = append(prompts, spec)
	}
	prompts = append(prompts, genai.Text("Create golang grpc table driven tests based on the above proto."))

	ct := time.Now().Format("2006-01-02-15-04-05")
	files.CheckDirectryExists()
	outputFile, err := os.Create(fmt.Sprintf("output/%s_output_test.md", ct))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	response := c.ProModel().GenerateContentStream(ctx, prompts...)
	c.CountTokens(ctx, prompts)
	for {
		resp, err := response.Next()
		if err == iterator.Done {
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
				fmt.Fprintln(writer, c)
			}
		}
	}
	return nil
}
