package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/posts"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/user"
)

const (
	port = 8081

	userServiceURL   = "http://user_service:8082"
	postsServiceHost = "localhost:8083"
)

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
	userHandlers := user.NewGateWayUserHandlers(userServiceURL)
	postsServiceHandlers, err := posts.NewGateWayPostsHandlers(userServiceURL, postsServiceHost)
	if err != nil {
		log.Fatalf("failed to start posts service handlers: %v", err)
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /signup", userHandlers.Proxy)
	router.HandleFunc("POST /signin", userHandlers.Proxy)
	router.HandleFunc("GET /profile", userHandlers.Proxy)
	router.HandleFunc("POST /profile", userHandlers.Proxy)

	router.HandleFunc("POST /posts/add", postsServiceHandlers.AddPost)
	router.HandleFunc("GET /posts/get", postsServiceHandlers.GetPost)
	router.HandleFunc("GET /posts/get_all", postsServiceHandlers.GetAllPosts)
	router.HandleFunc("GET /posts/get_all_of_author", postsServiceHandlers.GetAllPostsOfOneAuthor)
	router.HandleFunc("POST /posts/update", postsServiceHandlers.UpdatePost)
	router.HandleFunc("GET /posts/delete", postsServiceHandlers.DeletePost)

	handler := enableCORS(router)

	log.Println("Starting proxy-server on port", port)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalln(err)
	}
}
