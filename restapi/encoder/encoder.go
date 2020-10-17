package encoder

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/go-redis/redis"

	"github.com/krystalmejia24/samwise"
	"github.com/krystalmejia24/samwise/db"
	"github.com/krystalmejia24/samwise/encoder"
)

// Handler handles requests around encoders
type Handler struct {
	svc     *samwise.Service
	encoder encoder.Service
}

// New creates and returns a new Handler instance
func New(svc *samwise.Service) *Handler {
	return &Handler{svc: svc}
}

// Handle is to be used as a http.HandleFunc
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getEncoder(w, r)
	case http.MethodPost:
		h.postEncoder(w, r)
	case http.MethodDelete:
		h.deleteEncoder(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("method not supported"))
	}
}

func (h *Handler) getEncoder(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)

	encoder, err := db.GetEncoder(key, h.svc.DBConn)
	if err == redis.Nil {
		formatResponse(w, http.StatusNotFound, "No items with key "+key)
		return
	}
	if err != nil {
		formatResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	formatResponse(w, http.StatusOK, encoder.(string))
}

func (h *Handler) postEncoder(w http.ResponseWriter, r *http.Request) {
	var e db.Encoder

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		formatResponse(w, http.StatusBadRequest, "error marshalling the request body to an encoder")
		return
	}

	err = db.CreateEncoder(&e, h.svc.DBConn)
	if err != nil {
		formatResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	formatResponse(w, http.StatusOK, "Created encoder with key "+e.IP)
}

func (h *Handler) deleteEncoder(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.Path)

	val, err := db.DeleteEncoder(key, h.svc.DBConn)
	if err != nil {
		formatResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// no items were deleted
	if val == int64(0) {
		formatResponse(w, http.StatusNotFound, "No items deleted with key "+key)
		return
	}

	formatResponse(w, http.StatusOK, "Deleted encoder with key "+key)
}

func formatResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(msg))
}
