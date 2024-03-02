package gemini

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var apiKey = os.Getenv("GEMINI_API_KEY")

type Client interface {
	GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error)
	GenerateTextStream(ctx context.Context, specs []genai.Text, specType string) error
	GenerateTextStreamWriter(ctx context.Context, specs []genai.Text, language, specType string, outputFolder string) error
	GenerateVision(ctx context.Context, promptPart []genai.Part) (*genai.GenerateContentResponse, error)
	CompareImage(ctx context.Context, promptParts []string, search string) (*genai.GenerateContentResponse, error)
	GenerateVisionStream(ctx context.Context, prompt string) (*genai.GenerateContentResponseIterator, error)
	VisionModel() *genai.GenerativeModel
	ProModel() *genai.GenerativeModel
	Close()
}

type client struct {
	client *genai.Client
}

func New(ctx context.Context) (Client, error) {
	ai, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &client{
		client: ai,
	}, nil
}

func (c *client) Close() {
	c.client.Close()
}
