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

func (h *handler) RomDownload(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == domain.Empty {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(id); err != nil || strings.HasPrefix(id, "-") || strings.HasPrefix(id, "+") {
		http.Error(w, "Access Denied: id must be strictly numeric", http.StatusBadRequest)
		return
	}

	safedir := os.Getenv("ROM_STORAGE_PATH")

	searchPattern := filepath.Clean(filepath.Join(safedir, id, "*.7z"))

	if !strings.HasPrefix(searchPattern, safedir) {
		http.Error(w, "Access Denied: Invalid path sequence", http.StatusForbidden)
		return
	}

	matches, err := filepath.Glob(searchPattern)
	if err != nil || len(matches) != 1 {
		http.Error(w, "ROM file archive not found", http.StatusNotFound)
		return
	}

	file := matches[0]

	fileinfo, err := os.Stat(file)
	if err != nil || fileinfo.IsDir() {
		http.Error(w, "ROM file archive is unreadable", http.StatusNotFound)
		return
	}

	filesize := strconv.FormatInt(fileinfo.Size(), 10)
	romname := filepath.Base(file)

	w.Header().Set("Content-Disposition", "attachment; filename="+romname)
	w.Header().Set("Content-Type", "application/x-7z-compressed")
	w.Header().Set("Content-Length", filesize)

	http.ServeFile(w, r, file)
}
