package posts

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts: %v", err), http.StatusBadRequest)
		return
	}

	getAllPostsRequest := postspb.GetAllPostsRequest{}
	err = protojson.Unmarshal(body, &getAllPostsRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	getAllPostsResp, err := h.postsServiceClient.GetAllPosts(r.Context(), &getAllPostsRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getAllPostsResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("get all posts: error writing response: %v", err), http.StatusInternalServerError)
	}
}
