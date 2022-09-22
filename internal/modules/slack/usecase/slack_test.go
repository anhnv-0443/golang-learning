package usecase_test

import (
	"context"
	"go-app/internal/domain"
	"go-app/internal/modules/slack/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name string
	args *domain.SlackAppMentionEvent
	err  error
}

func TestHandleAppMentionEvent(t *testing.T) {
	t.Parallel()

	uc := usecase.NewUsecase()

	tcs := []TestCase{
		{
			name: "TODO",
			args: &domain.SlackAppMentionEvent{},
			err:  nil,
		},
	}

	for _, tc := range tcs {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := uc.HandleAppMentionEvent(context.Background(), tc.args)
			if err != nil {
				assert.ErrorIs(t, err, tc.err)
			}
		})
	}
}
