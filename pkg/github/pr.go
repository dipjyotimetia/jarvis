package github

import (
	"os"

	"github.com/google/go-github/v59/github"
)

var owner = os.Getenv("GITHUB_OWNER")
var repo = os.Getenv("GITHUB_REPO")

func (c *Client) ChekPR() (string, error) {
	pr, _, err := c.client.PullRequests.List(c.ctx, owner, repo,
		&github.PullRequestListOptions{
			State:       "open",
			Head:        "main",
			Sort:        "updated",
			Direction:   "desc",
			ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		return "", err
	}
	if len(pr) == 0 {
		return "", err
	}
	for _, p := range pr {
		ok, err := checkPrReference(p)
		if err != nil {
			return "", err
		}
		if ok {
			prDetails, _, err := c.client.PullRequests.Get(c.ctx, owner, repo, p.GetNumber())
			if err != nil {
				return "", err
			}
			if prDetails.GetMerged() {
				return "", err
			}
			return prDetails.GetBody(), nil
		}
	}
	return "", err
}

func checkPrReference(pr *github.PullRequest) (bool, error) {
	switch pr.GetHead().GetRef() {
	case "main":
		return false, nil
	case "master":
		return false, nil
	case "develop":
		return true, nil
	case "release":
		return false, nil
	case "hotfix":
		return true, nil
	case "feature":
		return true, nil
	case "bugfix":
		return true, nil
	case "chore":
		return true, nil
	case "docs":
		return false, nil
	case "test":
		return true, nil
	default:
		return false, nil
	}
}
