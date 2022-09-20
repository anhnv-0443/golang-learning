package http

import "go-app/internal/domain"

// SlackRequest is request from Slack
// https://api.slack.com/apis/connections/events-api#the-events-api__receiving-events__callback-field-overview
type SlackRequest struct {
	Token          string               `json:"token"`
	TeamID         string               `json:"team_id"`
	APIAppID       string               `json:"api_app_id"`
	Event          SlackEventRequest    `json:"event"`
	Type           string               `json:"type"`
	AuthedUsers    []string             `json:"authed_users"`
	AuthedTeams    []string             `json:"authed_teams"`
	Authorizations []SlackAuthorization `json:"authorizations"`
	EventContext   string               `json:"event_context"`
	EventID        string               `json:"event_id"`
	EventTime      int                  `json:"event_time"`
	Challenge      string               `json:"challenge"`
}

// SlackEventRequest is request from Slack
// https://api.slack.com/apis/connections/events-api#the-events-api__receiving-events__event-type-structure
type SlackEventRequest struct {
	Type    string `json:"type"`
	EventTS string `json:"event_ts"`
	User    string `json:"user"`
	TS      string `json:"ts"`
	Item    string `json:"item"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

// SlackAuthorization is authorizations object
// https://api.slack.com/apis/connections/events-api#the-events-api__receiving-events__callback-field-overview
type SlackAuthorization struct {
	EnterpriseID string `json:"enterprise_id"`
	TeamID       string `json:"team_id"`
	UserID       string `json:"user_id"`
	IsBot        bool   `json:"is_bot"`
}

// SlackChallengeResponse is struct used for url_verification request
type SlackChallengeResponse struct {
	Challenge string `json:"challenge"`
}

// SlackResponse is struct used for slack request
type SlackResponse struct {
	Ok bool `json:"ok"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

// ConvertRequestToSlackAppMentionEvent DTO
func ConvertRequestToSlackAppMentionEvent(req *SlackEventRequest) *domain.SlackAppMentionEvent {
	return &domain.SlackAppMentionEvent{
		Type:    req.Type,
		Channel: req.Channel,
		User:    req.User,
		Text:    req.Text,
		TS:      req.TS,
		EventTS: req.EventTS,
	}
}
