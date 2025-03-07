package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func (h *UserHandlers) SignIn(w http.ResponseWriter, req *http.Request) {
	body, err := readBodyFromRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	credentials := dto.LoginData{}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "error unmarshalling body: %v", err)
		return
	}

	userId, err := h.useCase.SignIn(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "error signing in: %v", err)
		return
	}

	err = h.setJWT(w, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}
}
