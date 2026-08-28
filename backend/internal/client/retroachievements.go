package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"retro-gallery/internal/config"
	"strings"
)

const (
	paramSeparator = "&"
)

type RetroAchievementsClient interface {
	GetApiKey() string
	GetUser() string
	Fetch(ctx context.Context, endpoint string, queryParams ...string) ([]byte, int, error)
}

type retroAchievementsClient struct {
	httpClient http.Client
	api        string
	username   string
	apiKey     string
}

func NewClient(username, apiKey string) *retroAchievementsClient {
	return &retroAchievementsClient{
		username: username,
		api:      config.RetroachievementsAPI,
		apiKey:   apiKey,
		httpClient: http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:        config.MaxIddleConns,
				IdleConnTimeout:     config.IddleConnTimeout,
				MaxIdleConnsPerHost: config.MaxIddleConns,
			},
		},
	}
}

func (c *retroAchievementsClient) GetApiKey() string {
	return c.apiKey
}

func (c *retroAchievementsClient) GetUser() string {
	return c.username
}

func (c *retroAchievementsClient) Fetch(ctx context.Context, endpoint string, queryParams ...string) ([]byte, int, error) {

	baseQuery := fmt.Sprintf("?y=%s&u=%s", c.apiKey, c.username)

	if len(queryParams) > 0 && queryParams[0] != "" {
		params := strings.Join(queryParams, "&")
		baseQuery = fmt.Sprintf("%s&%s", baseQuery, params)
	}

	url := fmt.Sprintf("%s%s%s", c.api, endpoint, baseQuery)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("internal server error: %w", err)
	}

	req.Header.Set("User-Agent", config.UserAgent)
	req.Header.Set("Accept", config.Accept)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, http.StatusBadGateway, fmt.Errorf("bad gateway: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("internal server error: %w", err)
	}

	return body, resp.StatusCode, nil

}
