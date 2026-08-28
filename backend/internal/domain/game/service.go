package game

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"retro-gallery/internal/cache"
	"retro-gallery/internal/client"
	"retro-gallery/internal/domain"
	"strings"
	"time"
)

type service struct {
	client client.RetroAchievementsClient
	cache  cache.Cache
}

const (
	getUserCompletedGamesPath      = "/API_GetUserCompletedGames.php"
	getGameInfoAndUserProgressPath = "/API_GetGameInfoAndUserProgress.php"
	completedGamesTTL              = 30 * time.Minute
	gameDetailsTTL                 = 5 * time.Minute
)

func NewService(c client.RetroAchievementsClient, ch cache.Cache) *service {
	return &service{
		client: c,
		cache:  ch,
	}
}

func (s *service) fetchCached(ctx context.Context, endpoint, user string, ttl time.Duration, params ...string) (data []byte, statusCode int, hit bool, err error) {
	var key string
	var query string

	if len(params) > 0 && params[0] != domain.Empty {
		query = strings.Join(params, domain.ParamSeparator)
		key = fmt.Sprintf("%s?u=%s?&%s", endpoint, user, query)
	} else {
		key = fmt.Sprintf("%s?u=%s", endpoint, user)
	}

	if data, hit := s.cache.Get(key); hit {
		return data, http.StatusOK, true, nil
	}

	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	data, statusCode, err = s.client.Fetch(cctx, endpoint, query)
	if err != nil {
		return nil, statusCode, false, err
	}

	if statusCode == http.StatusOK {
		s.cache.Set(key, data, ttl)
	}

	return data, statusCode, false, nil
}

func (s *service) GetCompletedGames(ctx context.Context, user string) (r *[]GameProgress, statusCode int, hit bool, err error) {
	raw, statusCode, hit, err := s.fetchCached(
		ctx, getUserCompletedGamesPath, user, completedGamesTTL,
	)
	if err != nil {
		return nil, statusCode, hit, err
	}

	var gameProgress []GameProgress
	if err := json.Unmarshal(raw, &gameProgress); err != nil {
		return nil, http.StatusInternalServerError, hit, err
	}

	return &gameProgress, statusCode, hit, nil
}

func (s *service) GetGameDetails(ctx context.Context, user, gameID string) (r *DetailedGameProgress, statusCode int, hit bool, err error) {
	raw, statusCode, hit, err := s.fetchCached(
		ctx, getGameInfoAndUserProgressPath, user, gameDetailsTTL,
		fmt.Sprintf("g=%s", gameID),
	)
	if err != nil {
		return nil, statusCode, hit, err
	}

	var detailedDetailed DetailedGameProgress
	if err := json.Unmarshal(raw, &detailedDetailed); err != nil {
		return nil, http.StatusInternalServerError, hit, err
	}

	return &detailedDetailed, statusCode, hit, nil
}
