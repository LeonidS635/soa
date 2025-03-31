package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/services/user/dto"
)

func (h *UserHandlers) SignUp(w http.ResponseWriter, req *http.Request) {
	body, err := readBodyFromRequest(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	credentials := dto.RegistrationData{}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "error unmarshalling body: %v", err)
		return
	}

	userId, err := h.useCase.SignUp(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		_, _ = fmt.Fprintf(w, "error signing up: %v", err)
		return
	}

	err = h.setJWT(w, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "%v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
