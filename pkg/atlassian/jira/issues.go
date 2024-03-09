package jira

import (
	"context"
	"fmt"
	"log"
)

func (c *client) GetIssues() {
	issues, _, err := c.Client.Issue.Get(context.Background(), "JAR", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(issues.Fields.Description)

}
