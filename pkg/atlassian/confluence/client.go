package confluence

import (
	"context"
	"log"
	"os"

	confluence "github.com/ctreminiom/go-atlassian/confluence/v2"
)

var (
	HOST     = os.Getenv("CONFLUENCE_HOST")
	USER     = os.Getenv("CONFLUENCE_USER")
	ApiToken = os.Getenv("CONFLUENCE_API_TOKEN")
)

type client struct {
	*confluence.Client
}

type Client interface {
	GetPages() ([]string, error)
}

func New(ctx context.Context) *client {
	confluenceClient, err := confluence.New(nil, HOST)
	if err != nil {
		log.Fatal(err)
	}

	confluenceClient.Auth.SetBasicAuth(USER, ApiToken)

	return &client{confluenceClient}
}

func (c *client) GetPages() ([]string, error) {
	pages, _, err := c.Client.Page.Gets(context.Background(), nil, "", 0)
	if err != nil {
		return nil, err
	}

	var pageTitles []string
	for _, page := range pages.Results {
		pageTitles = append(pageTitles, page.Title)
	}

	return pageTitles, nil
}
