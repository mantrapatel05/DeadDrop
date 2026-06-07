package server

import (
	"deadrop/crypto"
	"encoding/base64"
	"errors"
	"net/http"
)

func (s *Server) getDrop(w http.ResponseWriter, r *http.Request) {
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
		writeError(w, http.StatusGone, "drop is burned")
		return
	}

	if !drop.IsReady() {
		writeError(w, http.StatusTooEarly, "drop is not ready")
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(drop.Ciphertext)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to decode ciphertext")
		return
	}

	nonce, err := base64.StdEncoding.DecodeString(drop.Nonce)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to decode nonce")
		return
	}

	plaintext, err := crypto.Decrypt(ciphertext, nonce, s.encryptionKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to decrypt content")
		return
	}

	if err := s.MarkBurned(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to mark drop as burned")
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"content": string(plaintext),
	})
}
