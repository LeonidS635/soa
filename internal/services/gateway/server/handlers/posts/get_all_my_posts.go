package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetAllMyPosts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if !q.Has("page") {
		http.Error(w, "get all my posts: page parameter is required", http.StatusBadRequest)
	}

	pageStr := q.Get("page")
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		http.Error(w, "get all my posts: page parameter must be an integer", http.StatusBadRequest)
	}

	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	getAllMyPostsRequest := postspb.GetAllMyPostsRequest{
		UserId: userId,
		Page:   int32(page),
	}
	getAllMyPostsResp, err := h.postsServiceClient.GetAllMyPosts(
		r.Context(), &getAllMyPostsRequest,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts of one author: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getAllMyPostsResp)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get all posts of one author: error marshalling response: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(
			w, fmt.Sprintf("get all posts of one author: error writing response: %v", err),
			http.StatusInternalServerError,
		)
	}
}
