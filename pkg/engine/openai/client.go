package openai

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type Client struct {
	Client *azopenai.Client
}

func New(ctx context.Context) (*Client, error) {
	keyCredential := azcore.NewKeyCredential("API_KEY")
	client, err := azopenai.NewClientForOpenAI("https://api.openai.com/v1", keyCredential, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: client,
	}, nil
}
