package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	ApiKey = os.Getenv("API_KEY")
)

// GenClient is a wrapper around the generative ai client
type GenClient struct {
	ctx    context.Context
	Client *genai.Client
}

func NewGenClient(ctx context.Context) (*GenClient, error) {
	ai, err := genai.NewClient(ctx, option.WithAPIKey(ApiKey))
	if err != nil {
		return nil, err
	}
	return &GenClient{ctx, ai}, nil
}

func (c *GenClient) Close() {
	c.Client.Close()
}

func (c *GenClient) ProModel() *genai.GenerativeModel {
	proModel := c.Client.GenerativeModel("gemini-1.0-pro")
	fmt.Println(c.Client.ListModels(c.ctx))
	proModel.SetTemperature(0.8)
	proModel.SetTopK(40)
	proModel.SetTopP(0.8)

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
