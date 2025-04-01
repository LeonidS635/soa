package handlers

import (
	"encoding/binary"
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

	err = binary.Write(w, binary.LittleEndian, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "error writing body: %v", err)
		return
	}
}
