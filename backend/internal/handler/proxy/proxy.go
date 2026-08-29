package proxy

import "net/http"

type Handler interface {
	GetPlayerSummary(w http.ResponseWriter, r *http.Request)
	GetRecentUnlocks(w http.ResponseWriter, r *http.Request)
	GetCompletedGames(w http.ResponseWriter, r *http.Request)
	GetGameDetails(w http.ResponseWriter, r *http.Request)
}
