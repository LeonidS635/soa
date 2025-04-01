package posts

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("delete post: %v", err), http.StatusBadRequest)
		return
	}

	deletePostRequest := postspb.DeletePostRequest{
		UserId: userId,
	}
	err = protojson.Unmarshal(body, &deletePostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("delete post: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	deletePostResp, err := h.postsServiceClient.DeletePost(r.Context(), &deletePostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("delete post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(deletePostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("delete post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("delete post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
