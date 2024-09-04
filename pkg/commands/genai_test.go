package commands

import (
	"context"
	"testing"

	"github.com/dipjyotimetia/jarvis/pkg/atlassian/confluence"
	"github.com/dipjyotimetia/jarvis/pkg/atlassian/jira"
	"github.com/dipjyotimetia/jarvis/pkg/engine/gemini"
	"github.com/stretchr/testify/assert"
)

func TestReadConfluenceJira(t *testing.T) {
	ctx := context.Background()

	// Mock Confluence client
	confluenceClient := &confluence.MockClient{}
	confluenceClient.On("GetPages").Return([]string{"Page1", "Page2"}, nil)

	// Mock Jira client
	jiraClient := &jira.MockClient{}
	jiraClient.On("GetIssues").Return([]string{"Issue1", "Issue2"}, nil)

	// Mock Gemini client
	geminiClient := &gemini.MockClient{}
	geminiClient.On("GenerateTextStream", ctx, []string{"Page1", "Page2", "Issue1", "Issue2"}, "confluence-jira").Return(nil)

	// Call the function
	err := ReadConfluenceJira(ctx, confluenceClient, jiraClient, geminiClient)

	// Assert no error
	assert.NoError(t, err)

	// Assert that the mocks were called
	confluenceClient.AssertCalled(t, "GetPages")
	jiraClient.AssertCalled(t, "GetIssues")
	geminiClient.AssertCalled(t, "GenerateTextStream", ctx, []string{"Page1", "Page2", "Issue1", "Issue2"}, "confluence-jira")
}
