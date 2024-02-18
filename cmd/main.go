package main

import (
	"context"

	"github.com/dipjyotimetia/jarvis/pkg/engine/ai"
	"github.com/dipjyotimetia/jarvis/pkg/engine/files"
)

//find := os.Getenv("INPUT_FIND")

func main() {
	ctx := context.Background()
	// args := os.Args[0]
	ai, err := ai.NewGenClient(ctx)
	if err != nil {
		panic(err)
	}
	// fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))

	file, _ := files.IdentifySpecTypes("specs/proto")

	reader, _ := files.ReadFile(file[0])
	err = ai.GenerateTextStreamWriter(ctx, reader)
	if err != nil {
		panic(err)
	}

	// for _, v := range prompt.Candidates {
	// 	for _, vv := range v.Content.Parts {
	// 		fmt.Println(vv)
	// 	}
	// }
}
