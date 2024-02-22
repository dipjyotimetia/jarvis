package atlassian

import (
	"context"
	"log"

	jira "github.com/ctreminiom/go-atlassian/jira/v2"
)

type client struct {
	*jira.Client
}

func New(ctx context.Context) *client {

	jiraHost := "https://<_jira_instance_>.atlassian.net"
	mailAddress := "<your_mail>"
	apiToken := "<your_api_token>"

	jiraClient, err := jira.New(nil, jiraHost)
	if err != nil {
		log.Fatal(err)
	}

	jiraClient.Auth.SetBasicAuth(mailAddress, apiToken)

	return &client{jiraClient}
}
