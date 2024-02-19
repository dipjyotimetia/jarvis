package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
)

func (c *GenClient) CountTokens(ctx context.Context, parts []genai.Part) {
	resp, err := c.ProModel().CountTokens(ctx, parts...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.TotalTokens)
}
