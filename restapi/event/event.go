package event

import (
	"net/http"

	"github.com/krystalmejia24/samwise"
	"github.com/krystalmejia24/samwise/encoder"
	"github.com/krystalmejia24/samwise/event"
)

// Handler handles requests around encoders
type Handler struct {
	event   *event.Service
	encoder *encoder.Service
	svc     *samwise.Service
}

// New creates and returns a new Handler instance
func New(svc *samwise.Service) *Handler {
	return &Handler{svc: svc}
}

// Handle is to be used as a http.HandleFunc
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.postEvent(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("method not supported"))
	}
}

func (h *Handler) postEvent(w http.ResponseWriter, r *http.Request) {
	formatResponse(w, http.StatusOK, "TODO EVENT SERVICE")
}

func formatResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(msg))
}
