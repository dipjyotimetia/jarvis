package openai

import (
	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"os"
)

var apiKey = os.Getenv("OPEN_AI_API_KEY")

type Client interface {
	GenerateText() (string, error)
	GenerateTextStream() error
}

type client struct {
	client *azopenai.Client
}

func New() (Client, error) {
	keyCredential := azcore.NewKeyCredential(apiKey)
	ai, err := azopenai.NewClientForOpenAI("https://api.openai.com/v1", keyCredential, nil)
	if err != nil {
		return nil, err
	}
	return &client{
		ai,
	}, nil
}
