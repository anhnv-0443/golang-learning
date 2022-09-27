package usecase

import (
	"context"
	"go-app/internal/domain"
)

// SlackUsecase ...
type SlackUsecase struct{}

// NewUsecase will create new SlackUsecase object
func NewUsecase() domain.SlackUsecase {
	return &SlackUsecase{}
}

// HandleAppMentionEvent handle app_mention from slack
func (uc *SlackUsecase) HandleAppMentionEvent(ctx context.Context, event *domain.SlackAppMentionEvent) error {
	// TODO handle message to BOT

	return nil
}
