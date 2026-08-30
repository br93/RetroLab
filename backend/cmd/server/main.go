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
	"retro-gallery/internal/handler/download"
	"retro-gallery/internal/handler/proxy"
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
	proxyHandler := proxy.NewHandler(client, playerService, achievementService, gameService)
	downloadHandler := download.NewHandler()
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/retro/player", proxyHandler.GetPlayerSummary)
	mux.HandleFunc("/api/v1/retro/recent", proxyHandler.GetRecentUnlocks)
	mux.HandleFunc("/api/v1/retro/completed", proxyHandler.GetCompletedGames)
	mux.HandleFunc("/api/v1/retro/details", proxyHandler.GetGameDetails)
	mux.HandleFunc("/api/v1/download", downloadHandler.Download)

	serverAddr := config.ServerPort

	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Critical platform failure: %v", err)
	}
}
