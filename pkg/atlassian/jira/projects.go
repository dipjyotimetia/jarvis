package jira

import (
	"context"
	"fmt"
	"log"
)

func (c *client) GetProjects() {
	project, _, err := c.Client.Project.Get(context.Background(), "JAR", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(project.Key)

}
