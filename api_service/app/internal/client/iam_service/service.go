package iam_service

import (
	"context"
	"net/http"
	"time"

	"github.com/dexises/4room/api_service/pkg/logging"
	"github.com/dexises/4room/api_service/pkg/rest"
)

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, resource string, logger logging.Logger) IAMService {
	return &client{
		Resource: resource,
		base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logger,
		},
	}
}

type IAMService interface {
	GetUserToken(ctx context.Context, userID int) ([]byte, error)
	CreateToken(ctx context.Context, userID int) (string, error)
}

func (c *client) GetUserToken(ctx context.Context, userID int) ([]byte, error) {
	return []byte(""), nil
}

func (c *client) CreateToken(ctx context.Context, userID int) (string, error) {
	return "", nil
}
