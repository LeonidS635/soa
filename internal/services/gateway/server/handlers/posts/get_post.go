package posts

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetPost(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post: %v", err), http.StatusBadRequest)
		return
	}

	getPostRequest := postspb.GetPostRequest{
		UserId: userId,
	}
	err = protojson.Unmarshal(body, &getPostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	getPostResp, err := h.postsServiceClient.GetPost(r.Context(), &getPostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getPostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("get post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
