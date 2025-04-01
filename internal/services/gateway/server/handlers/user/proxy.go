package user

import (
	"io"
	"net/http"
)

func (h *GateWayUserHandlers) Proxy(w http.ResponseWriter, r *http.Request) {
	proxyReq, err := http.NewRequest(r.Method, h.userServiceURL+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error copying response body: "+err.Error(), http.StatusInternalServerError)
	}
}
