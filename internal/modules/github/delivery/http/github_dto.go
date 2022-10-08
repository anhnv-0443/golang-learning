package http

// GithubConnectRequest is request from user to connect github
type GithubConnectRequest struct {
	ClientId string `json:"client_id"`
}

// GithubCallbackRequest is request from github to server
type GithubCallbackRequest struct {
	Code string `json:"code"`
}

// ErrorResponse is struct when error
type ErrorResponse struct {
	Message string `json:"message"`
}

type GithubResponse struct {
	Ok bool `json:"ok"`
}
