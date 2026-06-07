package server

import (
	"errors"
	"net/http"
)

func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	drop, err := s.GetDrop(r.Context(), id)
	if err != nil {
		if errors.Is(err, errNotFound) {
			writeError(w, http.StatusNotFound, "drop not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to fetch drop")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ready":        drop.IsReady(),
		"burned":       drop.IsBurned(),
		"knock_count":  drop.KnockCount,
		"knock_target": drop.KnockTarget,
	})
}
