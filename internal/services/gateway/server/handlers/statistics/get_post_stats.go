package stats

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayStatisticsHandlers) GetPostStats(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.PathValue("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post statistics: error parsing post_id: %v", err), http.StatusBadRequest)
		return
	}

	getPostStatsRequest := statspb.GetPostStatsRequest{
		PostId: int32(postID),
	}
	getPostStatsResp, err := h.statisticsServiceClient.GetPostStats(r.Context(), &getPostStatsRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post statistics: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getPostStatsResp)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get post statistics: error marshalling response: %v", err), http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(
			w, fmt.Sprintf("get post statistics: error writing response: %v", err), http.StatusInternalServerError,
		)
	}
}
