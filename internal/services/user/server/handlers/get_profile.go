package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *UserHandlers) GetProfile(w http.ResponseWriter, req *http.Request) {
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

	profile, err := h.useCase.GetProfile(userId)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "error getting profile: %v", err)
		return
	}

	body, err := json.Marshal(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "error marshalling body: %v", err)
		return
	}

	written, err := w.Write(body)
	if err != nil || written != len(body) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "error writing body: %v", err)
		return
	}
}
