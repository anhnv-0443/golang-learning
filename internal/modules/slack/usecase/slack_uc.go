package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"go-app/config"
	"go-app/internal/domain"
	"io/ioutil"
	"net/http"
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
	_, err := uc.SendMessage(ctx, event.Channel, "Nice to meet you", event.User)
	if err != nil {
		return err
	}

	return nil
}

// SendMessage handle send message to slack
func (uc *SlackUsecase) SendMessage(
	ctx context.Context,
	channel string,
	content string,
	mentionUser string,
) (string, error) {
	content = "<@" + mentionUser + ">" + content
	values := map[string]string{"channel": channel, "text": content}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return "", nil
	}

	slackConfig := config.GetSlackConfig()
	req, err := http.NewRequest(http.MethodPost, slackConfig.SlackAPI, bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", nil
	}

	req.Header.Add("Authorization", "Bearer "+slackConfig.SlackToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
