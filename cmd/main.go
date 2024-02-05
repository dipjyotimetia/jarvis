package main

import (
	"context"
	"fmt"
	"jarvis/pkg/engine"
)

func main() {
	ctx := context.Background()
	ai, err := engine.NewGenClient(ctx)
	if err != nil {
		panic(err)
	}
	prompt, err := ai.GenerateText(ctx, "hello Gemini")
	if err != nil {
		panic(err)
	}

	for _, v := range prompt.Candidates {
		for _, vv := range v.Content.Parts {
			fmt.Println(vv)
		}
	}
}
