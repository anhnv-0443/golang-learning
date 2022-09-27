package domain

import (
	"context"
)

// SlackAppMentionEvent entity
// https://api.slack.com/events/app_mention
type SlackAppMentionEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// SlackUsecase represent the user's repository contract
type SlackUsecase interface {
	HandleAppMentionEvent(context.Context, *SlackAppMentionEvent) error
}
