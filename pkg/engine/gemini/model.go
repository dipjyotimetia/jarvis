package gemini

import "github.com/google/generative-ai-go/genai"

const (
	GEMINI_PRO        = "gemini-1.0-pro"
	GEMINI_1_5_PRO    = "gemini-1.5-pro"
	GEMINI_PRO_VISION = "gemini-pro-vision"
	EMBEDDING         = "embedding-001"
)

// Model returns the model
func (c *client) ProModel() *genai.GenerativeModel {
	proModel := c.client.GenerativeModel(GEMINI_PRO)
	proModel.SetTemperature(0.1)
	proModel.SetTopK(40)
	proModel.SetTopP(0.9)

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

// Model returns the model
func (c *client) ProFileModel() *genai.GenerativeModel {
	proModel := c.client.GenerativeModel(GEMINI_1_5_PRO)
	proModel.SetTemperature(0.1)
	proModel.SetTopK(40)
	proModel.SetTopP(0.9)

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

// VisionModel returns the vision model
func (c *client) VisionModel() *genai.GenerativeModel {
	visionModel := c.client.GenerativeModel(GEMINI_PRO_VISION)
	visionModel.SetTemperature(0.1)
	visionModel.SetTopK(40)
	visionModel.SetTopP(0.9)

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
