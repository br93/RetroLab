package player

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

const (
	getUserSummaryPath = "/API_GetUserSummary.php"
	playerSummaryTTL   = 15 * time.Minute
)

type service struct {
	client client.RetroAchievementsClient
	cache  cache.Cache
}

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

func (s *service) GetPlayerSummary(ctx context.Context, user, recentGames, recentAchievements string) (r *PlayerResponse, statusCode int, hit bool, err error) {
	raw, statusCode, hit, err := s.fetchCached(
		ctx, getUserSummaryPath, user, playerSummaryTTL,
		fmt.Sprintf("g=%s", recentGames),
		fmt.Sprintf("a=%s", recentAchievements),
	)
	if err != nil {
		return nil, statusCode, hit, err
	}

	var userProfile UserProfile
	if err := json.Unmarshal(raw, &userProfile); err != nil {
		return nil, http.StatusInternalServerError, hit, err
	}

	r = &PlayerResponse{
		User:            userProfile.User,
		RichPresenceMsg: userProfile.RichPresenceMsg,
		LastGameID:      userProfile.LastGameID,
		TotalPoints:     userProfile.TotalPoints,
		ID:              userProfile.ID,
		UserPic:         userProfile.UserPic,
	}

	return r, statusCode, hit, nil
}
