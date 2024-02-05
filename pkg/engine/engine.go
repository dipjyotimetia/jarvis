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

const (
	GEMINI_PRO        = "gemini-pro"
	GEMINI_PRO_VISION = "gemini-pro-vision"
)

type GenClient struct {
	Client *genai.Client
	Vision *genai.GenerativeModel
	Text   *genai.GenerativeModel
}

func NewGenClient(ctx context.Context) (*GenClient, error) {
	ai, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		return nil, err
	}
	visionModel := ai.GenerativeModel("gemini-pro-vision")
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

	proModel := ai.GenerativeModel("gemini-pro")
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
	return &GenClient{ai, visionModel, proModel}, nil
}

func (c *GenClient) Close() {
	c.Client.Close()
}

func (c *GenClient) GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	resp, err := c.Text.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}
	return resp, err
}

/*
imgData1, err := os.ReadFile(pathToImage1)

	if err != nil {
	  log.Fatal(err)
	}

imgData2, err := os.ReadFile(pathToImage1)

	if err != nil {
	  log.Fatal(err)
	}

	prompt := []genai.Part{
	  genai.ImageData("jpeg", imgData1),
	  genai.ImageData("jpeg", imgData2),
	  genai.Text("What's different between these two pictures?"),
	}
*/
func (c *GenClient) GenerateVision(ctx context.Context, promptPart []genai.Part) (*genai.GenerateContentResponse, error) {
	resp, err := c.Vision.GenerateContent(ctx, promptPart...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
