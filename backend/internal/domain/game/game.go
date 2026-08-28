package game

import "context"

type Service interface {
	GetCompletedGames(ctx context.Context, user string) (r *[]GameProgress, statusCode int, hit bool, err error)
	GetGameDetails(ctx context.Context, user, gameID string) (r *DetailedGameProgress, statusCode int, hit bool, err error)
}

type GameProgress struct {
	GameID       int    `json:"GameID"`
	Title        string `json:"Title"`
	ImageIcon    string `json:"ImageIcon"`
	ConsoleID    int    `json:"ConsoleID"`
	ConsoleName  string `json:"ConsoleName"`
	MaxPossible  int    `json:"MaxPossible"`
	NumAwarded   int    `json:"NumAwarded"`
	PctWon       string `json:"PctWon"`
	HardcoreMode string `json:"HardcoreMode"`
}

type DetailedGameProgress struct {
	ID              int                 `json:"ID"`
	Title           string              `json:"Title"`
	ConsoleName     string              `json:"ConsoleName"`
	ConsoleID       int                 `json:"ConsoleID"`
	ImageIcon       string              `json:"ImageIcon"`
	ImageTitle      string              `json:"ImageTitle"`
	ImageIngame     string              `json:"ImageIngame"`
	ImageBoxArt     string              `json:"ImageBoxArt"`
	NumAchievements int                 `json:"NumAchievements"`
	Achievements    map[string]RawBadge `json:"Achievements"`
	UserCompletion  string              `json:"UserCompletion"`
}

type RawBadge struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	BadgeName   string `json:"BadgeName"`
	Points      int    `json:"Points"`
	DateEarned  string `json:"DateEarned"`
}

type GameResponse struct {
	GameID      int    `json:"GameID"`
	Title       string `json:"Title"`
	ImageIcon   string `json:"ImageIcon"`
	ConsoleName string `json:"ConsoleName"`
	MaxPossible int    `json:"MaxPossible"`
	NumAwarded  int    `json:"NumAwarded"`
	PctWon      string `json:"PctWon"`
}
