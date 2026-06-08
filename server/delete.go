package server

import (
	"errors"
	"net/http"
)

func (s *Server) deleteDrop(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := s.MarkBurned(r.Context(), id); err != nil {
		if errors.Is(err, errNotFound) {
			writeError(w, http.StatusNotFound, "drop not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete drop")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}