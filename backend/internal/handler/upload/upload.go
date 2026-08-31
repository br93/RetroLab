package upload

import "net/http"

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}
