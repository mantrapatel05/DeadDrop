package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type patchRequest struct {
	RevealAt    *string `json:"reveal_at,omitempty"`
	KnockTarget *int    `json:"knock_target,omitempty"`
}

func (s *Server) patchDrop(w http.ResponseWriter, r *http.Request) {
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

	if drop.IsBurned() {
		writeError(w, http.StatusGone, "cannot modify a burned drop")
		return
	}

	var req patchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	var revealAt *time.Time
	if req.RevealAt != nil {
		t, err := time.Parse(time.RFC3339, *req.RevealAt)
		if err != nil {
			writeError(w, http.StatusBadRequest, "reveal_at must be RFC3339 format")
			return
		}
		revealAt = &t
	}

	if err := s.UpdateReveal(r.Context(), id, revealAt, req.KnockTarget); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update drop")
		return
	}

	updatedDrop, err := s.GetDrop(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch updated drop")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"reveal_at":    updatedDrop.RevealAt,
		"knock_target": updatedDrop.KnockTarget,
		"knock_count":  updatedDrop.KnockCount,
		"burned":       updatedDrop.IsBurned(),
		"ready":        updatedDrop.IsReady(),
	})
}
