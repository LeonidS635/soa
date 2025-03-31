package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func (h *UserHandlers) UpdateProfile(w http.ResponseWriter, req *http.Request) {
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

	body, err := readBodyFromRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	profile := dto.Profile{}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "error unmarshalling body: %v", err)
		return
	}

	err = h.useCase.UpdateProfile(userId, &profile)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "error updating profile: %v", err)
		return
	}
}
