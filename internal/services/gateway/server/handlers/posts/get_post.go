package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetPost(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	postIdStr := r.PathValue("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post: error parsing post_id: %v", err), http.StatusBadRequest)
		return
	}

	getPostRequest := postspb.GetPostRequest{
		UserId: userId,
		PostId: int32(postId),
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
