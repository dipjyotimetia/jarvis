package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func (c *GenClient) GenerateVision(ctx context.Context, promptPart []genai.Part) (*genai.GenerateContentResponse, error) {
	resp, err := c.VisionModel().GenerateContent(ctx, promptPart...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *GenClient) CompareImage(ctx context.Context, promptParts []string, search string) (*genai.GenerateContentResponse, error) {
	if len(promptParts) == 0 {
		return nil, fmt.Errorf("promptParts is empty")
	}
	if len(promptParts) == 1 {
		return c.GenerateVision(ctx, []genai.Part{genai.Text(promptParts[0]), genai.Text(search)})
	}
	if len(promptParts) > 2 {
		return nil, fmt.Errorf("promptParts is too long")
	}

	for _, p := range promptParts {
		imageData, err := os.ReadFile(p)
		if err != nil {
			return nil, err
		}
		promptParts = append(promptParts, string(imageData))
	}

	promptPart := []genai.Part{}
	for _, p := range promptParts {
		promptPart = append(promptPart, genai.ImageData("jpeg", []byte(p)), genai.Text(search))
	}

	resp, err := c.VisionModel().GenerateContent(ctx, promptPart...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *GenClient) GenerateVisionStream(ctx context.Context, prompt string) (*genai.GenerateContentResponseIterator, error) {
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
