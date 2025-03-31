package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	port           = 8081
	userServiceURL = "http://user_service:8082"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	proxyReq, err := http.NewRequest(r.Method, userServiceURL+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}
	proxyReq.ContentLength = r.ContentLength

	resp, err := http.DefaultTransport.RoundTrip(proxyReq)
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

func enableCORS(h *http.ServeMux) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			h.ServeHTTP(w, r)
		},
	)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", proxyHandler)
	handler := enableCORS(router)

	log.Println("Starting proxy-server on port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalln(err)
	}
}
