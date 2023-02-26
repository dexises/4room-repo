package user_service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dexises/4room/api_service/internal/apperror"
	"github.com/dexises/4room/api_service/pkg/logging"
	"github.com/dexises/4room/api_service/pkg/rest"
)

var _ UserService = &client{}

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, resource string, logger logging.Logger) UserService {
	c := client{
		Resource: resource,
		base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logger,
		},
	}
	return &c
}

type UserService interface {
	GetByEmailAndPassword(ctx context.Context, email, password string) (User, error)
	GetByID(ctx context.Context, id int) (User, error)
	Create(ctx context.Context, dto CreateUserDTO) (User, error)
	Update(ctx context.Context, id int, dto UpdateUserDTO) error
	Delete(ctx context.Context, id int) error
}

func (c *client) GetByEmailAndPassword(ctx context.Context, email, password string) (u User, err error) {
	filters := []rest.FilterOptions{
		{
			Field:  "email",
			Values: []string{email},
		},
		{
			Field:  "password",
			Values: []string{password},
		},
	}

	uri, err := c.base.BuildURL(c.Resource, filters)
	if err != nil {
		return u, fmt.Errorf("failed to build URL. error: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return u, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := c.base.SendRequest(req)
	if err != nil {
		return u, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		if err = json.NewDecoder(response.Body()).Decode(&u); err != nil {
			return u, fmt.Errorf("failed to decode body to error: %w", err)
		}
		return u, nil
	}

	return u, apperror.APIError(response.Error.Message, response.Error.DeveloperMessage, response.Error.ErrorCode)
}

func (c *client) GetByID(ctx context.Context, id int) (User, error) {
	return User{}, nil
}

func (c *client) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	return User{}, nil
}

func (c *client) Update(ctx context.Context, id int, dto UpdateUserDTO) error {
	return nil
}

func (c *client) Delete(ctx context.Context, id int) error {
	return nil
}
