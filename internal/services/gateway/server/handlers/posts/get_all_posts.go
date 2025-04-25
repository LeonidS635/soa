package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if !q.Has("page") {
		http.Error(w, "get all posts: page parameter is required", http.StatusBadRequest)
		return
	}

	pageStr := q.Get("page")
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		http.Error(w, "get all posts: invalid page number", http.StatusBadRequest)
		return
	}

	getAllPostsRequest := postspb.GetAllPostsRequest{
		Page: int32(page),
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
