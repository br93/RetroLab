package achievement

import "context"

type Service interface {
	GetRecentUnlocks(ctx context.Context, user, minutesLookBack string) (r *[]AchievementActivity, statusCode int, hit bool, err error)
}

type AchievementActivity struct {
	Date          string      `json:"Date"`
	HardcoreMode  int         `json:"HardcoreMode"`
	AchievementID int         `json:"AchievementID"`
	Title         string      `json:"Title"`
	Description   string      `json:"Description"`
	BadgeName     string      `json:"BadgeName"`
	Points        int         `json:"Points"`
	TrueRatio     int         `json:"TrueRatio"`
	Type          interface{} `json:"Type"`
	Author        string      `json:"Author"`
	AuthorULID    string      `json:"AuthorULID"`
	GameTitle     string      `json:"GameTitle"`
	GameIcon      string      `json:"GameIcon"`
	GameID        int         `json:"GameID"`
	ConsoleName   string      `json:"ConsoleName"`
	BadgeURL      string      `json:"BadgeURL"`
	GameURL       string      `json:"GameURL"`
}

type AchievementResponse struct {
	Date          string `json:"Date"`
	AchievementID int    `json:"AchievementID"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	BadgeName     string `json:"BadgeName"`
	Points        int    `json:"Points"`
	GameTitle     string `json:"GameTitle"`
	BadgeURL      string `json:"BadgeURL"`
	GameURL       string `json:"GameURL"`
}
