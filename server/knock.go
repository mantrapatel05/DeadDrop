package server

import(
	"errors"
	"net/http"
)

func (s *Server) knock(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	drop, err := s.GetDrop(r.Context(), id)

	if err != nil {
		if errors.Is(err, errNotFound){
			writeError(w, http.StatusNotFound, "drop not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to fetch drop")
		return
	}

	if drop.IsBurned(){
		writeError(w, http.StatusGone, "drop is burned")
		return
	}

	if drop.KnockTarget == nil{
		writeError(w, http.StatusBadRequest, "this is using time based unlock not knocks")
		return
	}

	if err := s.IncrementDrop(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to increment knock count")
		return
	}

	updatedDrop, err := s.GetDrop(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch the drop")
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"knock_count" : updatedDrop.KnockCount,
		"knock_target": updatedDrop.KnockTarget,
		"ready": updatedDrop.IsReady(),
	})
}