package domain

import (
	"context"
)

type GitHubGetCodeEvent struct {
	Code string `json:"code"`
}

// SlackUsecase represent the user's repository contract
type GithubUsecase interface {
	RequestCodeToGithub(ctx context.Context, clientId string, scope string) error
	GitHubCallbackGetCode(context.Context, *GitHubGetCodeEvent) error
	GetStatusPullRequest(ctx context.Context, owner string, repo string, pullNumber string) (string, error)
}
