package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) CommentPost(w http.ResponseWriter, r *http.Request) {
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

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("comment post: %v", err), http.StatusBadRequest)
		return
	}

	commentPostRequest := postspb.CommentPostRequest{
		UserId:  userId,
		PostId:  int32(postId),
		Comment: &postspb.Comment{},
	}
	err = protojson.Unmarshal(body, commentPostRequest.Comment)
	if err != nil {
		http.Error(w, fmt.Sprintf("comment post: error unmarshalling body: %v", err), http.StatusBadRequest)
		return
	}

	commentPostResp, err := h.postsServiceClient.CommentPost(r.Context(), &commentPostRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("comment post: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(commentPostResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("comment post: error marshalling response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(w, fmt.Sprintf("comment post: error writing response: %v", err), http.StatusInternalServerError)
	}
}
