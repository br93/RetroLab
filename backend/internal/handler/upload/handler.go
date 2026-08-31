package upload

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"retro-gallery/internal/domain"
	"strconv"
	"strings"
)

type handler struct {
}

const (
	bufferSize       = 32 << 20
	fileMode         = 0755
	extension        = "srm"
	extensionPattern = "*.srm"
)

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if _, err := strconv.Atoi(id); err != nil || strings.HasPrefix(id, "-") || strings.HasPrefix(id, "+") {
		http.Error(w, "Access denied: id must be strictly numeric", http.StatusBadRequest)
		return
	}

	safedir := os.Getenv("SAVE_STORAGE_PATH")
	if safedir == domain.Empty {
		http.Error(w, "Server configuration error: path invalid", http.StatusInternalServerError)
		return
	}

	targetdir := filepath.Clean(filepath.Join(safedir, id))

	err := r.ParseMultipartForm(bufferSize)
	if err != nil {
		http.Error(w, "Failed parsing form data", http.StatusBadRequest)
		return
	}

	uploaded, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Missing file in form", http.StatusBadRequest)
		return
	}
	defer uploaded.Close()

	if !strings.HasSuffix(strings.ToLower(header.Filename), extension) {
		http.Error(w, "Acesse Denied: invalid extension", http.StatusBadRequest)
	}

	if err := os.MkdirAll(targetdir, fileMode); err != nil {
		http.Error(w, "Failed to initialize", http.StatusInternalServerError)
	}

	existingsave, _ := filepath.Glob(filepath.Join(targetdir, extensionPattern))
	for _, f := range existingsave {
		_ = os.Remove(f)
	}

	destpath := filepath.Join(targetdir, filepath.Base(header.Filename))

	out, err := os.Create(destpath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, uploaded); err != nil {
		http.Error(w, "Failed to write file content", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success: File uploaded"))
}
