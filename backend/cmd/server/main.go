package main

import (
	"log"
	"net/http"
	"os"
	"retro-gallery/internal/cache"
	"retro-gallery/internal/client"
	"retro-gallery/internal/config"
	"retro-gallery/internal/domain/achievement"
	"retro-gallery/internal/domain/game"
	"retro-gallery/internal/domain/player"
	"retro-gallery/internal/handler"
)

func main() {

	apiKey := os.Getenv("RETRO_API_KEY")
	username := os.Getenv("RETRO_API_USERNAME")

	if apiKey == "" || username == "" {
		log.Fatal("Failed: Missing RETRO_API_KEY or RETRO_API_USERNAME")
	}

	cache := cache.NewCache()
	client := client.NewClient(username, apiKey)
	playerService := player.NewService(client, cache)
	achievementService := achievement.NewService(client, cache)
	gameService := game.NewService(client, cache)
	handler := handler.NewProxyHandler(client, playerService, achievementService, gameService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/retro/player", handler.GetPlayerSummary)
	mux.HandleFunc("/api/v1/retro/recent", handler.GetRecentUnlocks)
	mux.HandleFunc("/api/v1/retro/completed", handler.GetCompletedGames)
	mux.HandleFunc("/api/v1/retro/details", handler.GetGameDetails)

	serverAddr := config.ServerPort

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("❌ Critical platform failure: %v", err)
	}
}
