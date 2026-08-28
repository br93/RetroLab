package player

import (
	"context"
	"time"
)

type Service interface {
	GetPlayerSummary(ctx context.Context, user, recentGames, recentAchievements string) (r *PlayerResponse, statusCode int, hit bool, err error)
}

type UserProfile struct {
	User                string                 `json:"User"`
	MemberSince         string                 `json:"MemberSince"`
	LastActivity        Activity               `json:"LastActivity"`
	RichPresenceMsg     string                 `json:"RichPresenceMsg"`
	RichPresenceMsgDate string                 `json:"RichPresenceMsgDate"`
	LastGameID          int                    `json:"LastGameID"`
	ContribCount        int                    `json:"ContribCount"`
	ContribYield        int                    `json:"ContribYield"`
	TotalPoints         int                    `json:"TotalPoints"`
	TotalSoftcorePoints int                    `json:"TotalSoftcorePoints"`
	TotalTruePoints     int                    `json:"TotalTruePoints"`
	Permissions         int                    `json:"Permissions"`
	Untracked           int                    `json:"Untracked"`
	ID                  int                    `json:"ID"`
	UserWallActive      int                    `json:"UserWallActive"`
	Motto               string                 `json:"Motto"`
	Rank                int                    `json:"Rank"`
	RecentlyPlayedCount int                    `json:"RecentlyPlayedCount"`
	RecentlyPlayed      []interface{}          `json:"RecentlyPlayed"`
	ULID                string                 `json:"ULID"`
	UserPic             string                 `json:"UserPic"`
	TotalRanked         int                    `json:"TotalRanked"`
	Awarded             map[string]interface{} `json:"Awarded"`
	RecentAchievements  map[string]interface{} `json:"RecentAchievements"`
	Status              string                 `json:"Status"`
}

type Activity struct {
	ID           int         `json:"ID"`
	Timestamp    *time.Time  `json:"timestamp"`
	LastUpdate   *time.Time  `json:"lastupdate"`
	ActivityType interface{} `json:"activitytype"`
	User         string      `json:"User"`
	Data         interface{} `json:"data"`
	Data2        interface{} `json:"data2"`
}

type PlayerResponse struct {
	User            string `json:"User"`
	RichPresenceMsg string `json:"RichPresenceMsg"`
	LastGameID      int    `json:"LastGameID"`
	TotalPoints     int    `json:"TotalPoints"`
	ID              int    `json:"ID"`
	UserPic         string `json:"UserPic"`
}
