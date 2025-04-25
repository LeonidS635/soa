package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *UserHandlers) GetUserId(w http.ResponseWriter, req *http.Request) {
	token, err := h.getJWT(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	userId, err := getUserIdFromJWT(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	data, err := json.Marshal(
		struct {
			UserId int `json:"user_id"`
		}{UserId: userId},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "error marshalling body: %v", err)
		return
	}

	written, err := w.Write(data)
	if err != nil || written != len(data) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "error writing body: %v", err)
		return
	}
}
