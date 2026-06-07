package server

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"deadrop/crypto"
	"deadrop/model"

	"github.com/google/uuid"
)

type createRequest struct {
	Content     string  `json:"content"`
	RevealAt    *string `json:"reveal_at,omitempty"`
	KnockTarget *int    `json:"knock_target,omitempty"`
}

func (s *Server) createDrop(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Content == "" {
		writeError(w, http.StatusBadRequest, "content is required")
		return
	}

	var revealAt *time.Time
	if req.RevealAt != nil {
		t, err := time.Parse(time.RFC3339, *req.RevealAt)
		if err != nil {
			writeError(w, http.StatusBadRequest, "reveal_at must be RFC3339 format (e.g. 2026-06-07T00:00:00Z)")
			return
		}
		revealAt = &t
	}

	drop := model.Drop{
		RevealAt:    revealAt,
		KnockTarget: req.KnockTarget,
	}

	if err := drop.ValidCreate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	ciphertext, nonce, err := crypto.Encrypt([]byte(req.Content), s.encryptionKey)
	if err != nil {
		log.Printf("encryption error: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to encrypt content")
		return
	}

	id := uuid.New().String()
	now := time.Now()
	zero := 0
	falseVal := false

	drop.ID = id
	drop.Ciphertext = base64.StdEncoding.EncodeToString(ciphertext)
	drop.Nonce = base64.StdEncoding.EncodeToString(nonce)
	drop.CreatedAt = &now
	drop.Burned = &falseVal
	drop.KnockCount = &zero

	if err := s.CreateDrop(r.Context(), &drop); err != nil {
		log.Printf("database error: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create drop")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"id": id})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
