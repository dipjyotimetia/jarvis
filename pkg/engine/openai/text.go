package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func (c *client) GenerateText() (string, error) {

	resp, err := c.client.GetCompletions(context.TODO(), azopenai.CompletionsOptions{
		Prompt:         []string{"What is Azure OpenAI, in 20 words or less"},
		MaxTokens:      to.Ptr(int32(2048)),
		Temperature:    to.Ptr(float32(0.8)),
		DeploymentName: to.Ptr("gpt-3.5-turbo-instruct"),
	}, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no choices returned")
	}

	return *resp.Choices[0].Text, nil
}

func (c *client) GenerateTextStream() error {

	resp, err := c.client.GetCompletionsStream(context.TODO(), azopenai.CompletionsOptions{
		Prompt:         []string{"What is Azure OpenAI, in 20 words or less"},
		MaxTokens:      to.Ptr(int32(2048)),
		Temperature:    to.Ptr(float32(0.0)),
		DeploymentName: to.Ptr("gpt-35-turbo-16k"),
	}, nil)
	if err != nil {
		return err
	}

	for {
		completions, err := resp.CompletionsStream.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
		for _, choice := range completions.Choices {
			text := ""
			if choice.Text != nil {
				text = *choice.Text
			}
			fmt.Fprintf(os.Stderr, "Result: %s\n", text)
		}
	}
	return nil
}
