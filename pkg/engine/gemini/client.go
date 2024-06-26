package gemini

import (
	"context"
	"errors"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var apiKey = os.Getenv("GEMINI_API_KEY")

type Client interface {
	GenerateText(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error)
	GenerateTextStream(ctx context.Context, specs []genai.Text, specType string) error
	GenerateTextStreamFromFile(ctx context.Context, path string) error
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

// New creates a new Gemini client.
func New(ctx context.Context) (Client, error) {
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY is not set")
	}
	ai, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &client{
		client: ai,
	}, nil
}

// Close closes the Gemini client.
func (c *client) Close() {
	c.client.Close()
}
