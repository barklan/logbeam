package api

import (
	"io"
	"net/http"
)

type inputResponse struct {
	Accepted int `json:"accepted"`
}

func (s *Server) input(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		s.errorResp(w, r, "Failed to read body.", http.StatusBadRequest)

		return
	}
	w.WriteHeader(http.StatusAccepted)
}
