package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"retro-gallery/internal/client"
	"retro-gallery/internal/domain"
	"retro-gallery/internal/domain/achievement"
	"retro-gallery/internal/domain/game"
	"retro-gallery/internal/domain/player"
)

type handler struct {
	client             client.RetroAchievementsClient
	playerService      player.Service
	achievementService achievement.Service
	gameService        game.Service
}

const (
	recentGamesParam   = "g"
	recentGamesDefault = "0"

	recentAchievementsParam   = "a"
	recentAchievementsDefault = "10"

	minutesLookBackParam   = "m"
	minutesLookBackDefault = "43200"

	contentTypeHeader = "Content-Type"
	contentTypeValue  = "application/json"
	cacheHeader       = "X-Cache-Status"
)

func NewHandler(client client.RetroAchievementsClient, playerService player.Service, achievementService achievement.Service, gameService game.Service) *handler {
	return &handler{
		client:             client,
		playerService:      playerService,
		achievementService: achievementService,
		gameService:        gameService,
	}
}

func (h *handler) GetPlayerSummary(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	recentGames := query.Get(recentGamesParam)
	recentAchievements := query.Get(recentAchievementsParam)

	if recentGames == domain.Empty {
		recentGames = recentGamesDefault
	}

	if recentAchievements == domain.Empty {
		recentAchievements = recentAchievementsDefault
	}

	data, statusCode, hit, err := h.playerService.GetPlayerSummary(r.Context(), h.client.GetUser(), recentGames, recentAchievements)

	if err != nil {
		writeError(w, statusCode, err)
		return
	}

	body, err := json.Marshal(data)

	if hit {
		writeHit(w, body)
		return
	}
	writeMiss(w, body, statusCode)

}

func (h *handler) GetRecentUnlocks(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	minutesLookBack := query.Get(minutesLookBackParam)

	if minutesLookBack == domain.Empty {
		minutesLookBack = minutesLookBackDefault
	}

	data, statusCode, hit, err := h.achievementService.GetRecentUnlocks(r.Context(), h.client.GetUser(), minutesLookBack)

	if err != nil {
		writeError(w, statusCode, err)
		return
	}

	body, err := json.Marshal(data)

	if hit {
		writeHit(w, body)
		return
	}
	writeMiss(w, body, statusCode)
}

func (h *handler) GetCompletedGames(w http.ResponseWriter, r *http.Request) {
	data, statusCode, hit, err := h.gameService.GetCompletedGames(r.Context(), h.client.GetUser())

	if err != nil {
		writeError(w, statusCode, err)
		return
	}

	body, err := json.Marshal(data)

	if hit {
		writeHit(w, body)
		return
	}
	writeMiss(w, body, statusCode)
}

func writeHit(w http.ResponseWriter, data []byte) {
	w.Header().Set(contentTypeHeader, contentTypeValue)
	w.Header().Set(cacheHeader, "HIT")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func writeMiss(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set(contentTypeHeader, contentTypeValue)
	w.Header().Set(cacheHeader, "MISS")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set(contentTypeHeader, contentTypeValue)
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	return
}

func (h *handler) GetGameDetails(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	gameID := query.Get("id")

	if gameID == domain.Empty {
		writeError(w, http.StatusBadRequest, fmt.Errorf("missing query parameter: id"))
		return
	}

	data, statusCode, hit, err := h.gameService.GetGameDetails(r.Context(), h.client.GetUser(), gameID)
	if err != nil {
		writeError(w, statusCode, err)
		return
	}

	body, err := json.Marshal(data)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if hit {
		writeHit(w, body)
		return
	}
	writeMiss(w, body, statusCode)
}
