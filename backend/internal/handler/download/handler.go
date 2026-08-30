package download

import (
	"net/http"
	"os"
	"path/filepath"
	"retro-gallery/internal/domain"
	"strconv"
	"strings"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Download(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	rom := r.URL.Query().Get("rom")
	save := r.URL.Query().Get("save")

	if id == domain.Empty {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(id); err != nil || strings.HasPrefix(id, "-") || strings.HasPrefix(id, "+") {
		http.Error(w, "Access Denied: id must be strictly numeric", http.StatusBadRequest)
		return
	}

	var safedir string
	var extension string
	var contenttype string

	if rom == "true" {
		safedir = os.Getenv("ROM_STORAGE_PATH")
		extension = "*.7z"
		contenttype = "application/x-7z-compressed"
	} else if save == "true" {
		safedir = os.Getenv("SAVE_STORAGE_PATH")
		// ONLY SRM AT THE MOMENT, WILL IMPROVE LATER //
		extension = "*.srm"
		contenttype = "application/octet-stream"
	}

	if safedir == domain.Empty {
		http.Error(w, "Missing rom or save parameter", http.StatusBadRequest)
		return
	}

	searchPattern := filepath.Clean(filepath.Join(safedir, id, extension))

	if !strings.HasPrefix(searchPattern, safedir) {
		http.Error(w, "Access Denied: Invalid path sequence", http.StatusForbidden)
		return
	}

	matches, err := filepath.Glob(searchPattern)
	if err != nil || len(matches) != 1 {
		http.Error(w, "File archive not found", http.StatusNotFound)
		return
	}

	file := matches[0]

	fileinfo, err := os.Stat(file)
	if err != nil || fileinfo.IsDir() {
		http.Error(w, "File archive is unreadable", http.StatusNotFound)
		return
	}

	filesize := strconv.FormatInt(fileinfo.Size(), 10)
	filename := filepath.Base(file)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", contenttype)
	w.Header().Set("Content-Length", filesize)

	http.ServeFile(w, r, file)
}
