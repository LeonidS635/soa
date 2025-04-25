package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
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

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("update post: %v", err), http.StatusBadRequest)
		return
	}

	updatePostRequest := postspb.UpdatePostRequest{
		UserId: userId,
		PostId: int32(postId),
		Post:   &postspb.Post{},
	}
	err = protojson.Unmarshal(body, updatePostRequest.Post)
	if err != nil {
		http.Error(w, fmt.Sprintf("update post: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	updatePostResp, err := h.postsServiceClient.UpdatePost(r.Context(), &updatePostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("update post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(updatePostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("update post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("update post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
