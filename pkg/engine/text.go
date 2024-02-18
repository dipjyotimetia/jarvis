package engine

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

func (c *GenClient) GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	resp, err := c.ProModel().GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GenerateContentStream
func (c *GenClient) GenerateTextStream(ctx context.Context, prompt string) (*genai.GenerateContentResponseIterator, error) {
	resp := c.ProModel().GenerateContentStream(ctx, genai.Text(prompt))
	for {
		resp, err := resp.Next()
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
}
