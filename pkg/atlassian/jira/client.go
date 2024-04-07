package jira

import (
	"context"
	"log"
	"os"

	jira "github.com/ctreminiom/go-atlassian/jira/v2"
)

var (
	HOST     = os.Getenv("JIRA_HOST")
	USER     = os.Getenv("JIRA_USER")
	ApiToken = os.Getenv("JIRA_API_TOKEN")
)

type client struct {
	*jira.Client
}

type Client interface {
	GetIssues()
	GetProjects()
}

func New(ctx context.Context) *client {
	jiraClient, err := jira.New(nil, HOST)
	if err != nil {
		log.Fatal(err)
	}

	jiraClient.Auth.SetBasicAuth(USER, ApiToken)

	return &client{jiraClient}
}
