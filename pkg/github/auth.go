package github

import "github.com/google/go-github/v59/github"

type Client struct {
	client *github.Client
}

func NewClient(token string) *Client {
	return &Client{
		client: github.NewClient(nil),
	}
}
