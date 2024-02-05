package main

import (
	"context"
	"fmt"
	"jarvis/pkg/engine"
	"os"
)

func main() {
	ctx := context.Background()
	args := os.Args[0]
	ai, err := engine.NewGenClient(ctx)
	if err != nil {
		panic(err)
	}
	prompt, err := ai.GenerateText(ctx, args)
	if err != nil {
		panic(err)
	}

	for _, v := range prompt.Candidates {
		for _, vv := range v.Content.Parts {
			fmt.Println(vv)
		}
	}
}
