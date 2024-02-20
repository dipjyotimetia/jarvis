package openai

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

var (
	OpenAiApiKey = os.Getenv("OPEN_AI_API_KEY")
)

type GptClient struct {
	ctx    context.Context
	Client *azopenai.Client
}

func NewClient(ctx context.Context, apiKey string) *GptClient {
	keyCredential := azcore.NewKeyCredential(OpenAiApiKey)
	client, err := azopenai.NewClientForOpenAI("https://api.openai.com/v1", keyCredential, nil)
	if err != nil {
		panic(err)
	}
	return &GptClient{
		ctx:    ctx,
		Client: client,
	}
}
