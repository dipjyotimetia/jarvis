package gemini

import "github.com/google/generative-ai-go/genai"

func (c *client) ProModel() *genai.GenerativeModel {
	proModel := c.client.GenerativeModel("gemini-1.0-pro")
	proModel.SetTemperature(0.8)
	proModel.SetTopK(40)
	proModel.SetTopP(0.8)

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
	return proModel
}

func (c *client) VisionModel() *genai.GenerativeModel {
	visionModel := c.client.GenerativeModel("gemini-pro-vision")
	visionModel.SetTemperature(0.8)
	visionModel.SetTopK(40)
	visionModel.SetTopP(0.8)

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
	return visionModel
}
