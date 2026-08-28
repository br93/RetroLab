package config

import "time"

const (
	RetroachievementsAPI = "https://retroachievements.org/API"
	ServerPort           = ":5000"
	Timeout              = 10 * time.Second
	MaxIddleConns        = 100
	IddleConnTimeout     = 90 * time.Second
	MaxIdleConnsPerHost  = 20
	UserAgent            = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	Accept               = "application/json"
)
