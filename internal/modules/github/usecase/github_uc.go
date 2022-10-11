package usecase

import (
	"context"
	"go-app/internal/domain"
)

// SlackUsecase ...
type GithubUsecase struct{}

func (uc *GithubUsecase) RequestCodeToGithub(ctx context.Context, clientId string, scope string) error {
	return nil
}

func (uc *GithubUsecase) GitHubCallbackGetCode(ctx context.Context, event *domain.GitHubGetCodeEvent) error {
	return nil
}

func (uc *GithubUsecase) GetStatusPullRequest(ctx context.Context, owner string, repo string, pullNumber string) (string, error) {
	return "", nil
}

// NewUsecase will create new SlackUsecase object
func NewUsecase() domain.GithubUsecase {
	return &GithubUsecase{}
}
