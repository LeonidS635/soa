package posts

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/postspb"
	"github.com/LeonidS635/soa/internal/services/gateway/server/handlers/helpers"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayPostsHandlers) GetAllPostsOfOneAuthor(w http.ResponseWriter, r *http.Request) {
	userId, err := h.getUserId(w, r)
	if err != nil {
		return
	}

	body, err := helpers.ReadBodyFromRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts of one author: %v", err), http.StatusBadRequest)
		return
	}

	getAllPostsOfOneAuthorRequest := postspb.GetAllPostsOfOneAuthorRequest{
		UserId: userId,
	}
	err = protojson.Unmarshal(body, &getAllPostsOfOneAuthorRequest)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get all posts of one author: error unmarshalling body: %v", err), http.StatusBadRequest,
		)
		return
	}

	getAllPostsOfOneAuthorResp, err := h.postsServiceClient.GetAllPostsOfOneAuthor(
		r.Context(), &getAllPostsOfOneAuthorRequest,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("get all posts of one author: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getAllPostsOfOneAuthorResp)
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
