package download

import (
	"net/http"
)

type Handler interface {
	RomDownload(w http.ResponseWriter, r *http.Request)
}
