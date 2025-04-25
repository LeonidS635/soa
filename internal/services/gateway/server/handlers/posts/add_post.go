package posts

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) AddPost(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("add post: %v", err), http.StatusBadRequest)
		return
	}

	addPostRequest := postspb.AddPostRequest{
		UserId: userId,
		Post:   &postspb.Post{},
	}
	err = protojson.Unmarshal(body, addPostRequest.Post)
	if err != nil {
		http.Error(w, fmt.Sprintf("add post: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	addPostResp, err := h.postsServiceClient.AddPost(r.Context(), &addPostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("add post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(addPostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("add post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("add post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
