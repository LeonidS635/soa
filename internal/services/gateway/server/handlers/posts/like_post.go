package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) LikePost(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	postIdStr := r.PathValue("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("like post: error parsing post_id: %v", err), http.StatusBadRequest)
		return
	}

	likePostRequest := postspb.LikePostRequest{
		UserId: userId,
		PostId: int32(postId),
	}
	likePostResp, err := h.postsServiceClient.LikePost(r.Context(), &likePostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("like post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(likePostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("like post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("like post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
