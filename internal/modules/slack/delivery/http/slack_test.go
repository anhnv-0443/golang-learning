package http_test

import (
	"context"
	"errors"
	"go-app/internal/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockDomain "go-app/internal/domain/mock"
	slackHttp "go-app/internal/modules/slack/delivery/http"
)

var errMock = errors.New("mock Error")

type handleSlackRequestTestCase struct {
	name            string
	requestBody     string
	expectedCode    int
	expectedResBody string
	expectedError   error
	beforeHook      func(t *testing.T, req *http.Request, c echo.Context)
}

func TestHandleSlackRequest(t *testing.T) {
	e := echo.New()
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecaseMock := mockDomain.NewMockSlackUsecase(ctrl)
	handler := slackHttp.SlackHandler{Usecase: usecaseMock}

	tests := []handleSlackRequestTestCase{
		{
			name:         "UNPROCESSABLE ENTITY",
			requestBody:  `[]`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedResBody: "{\"message\":\"code=400, message=Unmarshal type error: expected=http.SlackRequest, " +
				"got=array, field=, offset=1, internal=json: cannot unmarshal array into Go value of type http.SlackRequest\"}\n",
		},
		{
			name:            "Event url_verification",
			requestBody:     `{"type": "url_verification", "challenge": "token"}`,
			expectedCode:    http.StatusOK,
			expectedResBody: "{\"challenge\":\"token\"}\n",
		},
		{
			name:            "Event app_mention",
			requestBody:     `{"type": "event_callback", "event": {"type": "app_mention"}}`,
			expectedCode:    http.StatusOK,
			expectedResBody: "{\"ok\":true}\n",
			beforeHook: func(t *testing.T, req *http.Request, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().HandleAppMentionEvent(context.Background(), &domain.SlackAppMentionEvent{Type: "app_mention"}).
					Times(1)
			},
		},
		{
			name:            "Event app_mention",
			requestBody:     `{"type": "event_callback", "event": {"type": "app_mention"}}`,
			expectedCode:    http.StatusOK,
			expectedResBody: "{\"ok\":false}\n",
			beforeHook: func(t *testing.T, req *http.Request, c echo.Context) {
				t.Helper()
				usecaseMock.EXPECT().HandleAppMentionEvent(context.Background(), &domain.SlackAppMentionEvent{Type: "app_mention"}).
					Times(1).Return(errMock)
			},
		},
		{
			name:            "Other Events",
			requestBody:     `{"type": "event_callback"}`,
			expectedCode:    http.StatusOK,
			expectedResBody: "{\"ok\":true}\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/slack/", strings.NewReader(tc.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tc.beforeHook != nil {
				tc.beforeHook(t, req, c)
			}

			err := handler.HandleSlackRequest(c)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Equal(t, tc.expectedResBody, rec.Body.String())
				assert.Equal(t, tc.expectedCode, rec.Code)
			}
		})
	}
}
