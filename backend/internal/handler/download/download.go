package download

import (
	"net/http"
)

type Handler interface {
	Download(w http.ResponseWriter, r *http.Request)
}
