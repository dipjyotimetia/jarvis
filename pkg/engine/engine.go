package engine

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	API_KEY = os.Getenv("API_KEY")
)

// GenClient is a wrapper around the generative ai client
type GenClient struct {
	Client *genai.Client
}

func NewGenClient(ctx context.Context) (*GenClient, error) {
	ai, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		return nil, err
	}
	return &GenClient{ai}, nil
}

func (c *GenClient) Close() {
	c.Client.Close()
}

func (c *GenClient) ProModel() *genai.GenerativeModel {
	proModel := c.Client.GenerativeModel("gemini-pro")

	proModel.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return proModel
}

func (c *GenClient) VisionModel() *genai.GenerativeModel {
	visionModel := c.Client.GenerativeModel("gemini-pro-vision")

	visionModel.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return visionModel
}
