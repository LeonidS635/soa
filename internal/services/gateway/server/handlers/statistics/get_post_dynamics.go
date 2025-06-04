package stats

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayStatisticsHandlers) GetPostDynamics(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.PathValue("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post dynamics: error parsing post_id: %v", err), http.StatusBadRequest)
		return
	}

	metricStr := r.URL.Query().Get("metric")
	var metric statspb.Metric
	switch metricStr {
	case metricViews:
		metric = statspb.Metric_VIEWS
	case metricLikes:
		metric = statspb.Metric_LIKES
	case metricComments:
		metric = statspb.Metric_COMMENTS
	default:
		http.Error(w, fmt.Sprintf("get post dynamics: unknown metric: %s", metricStr), http.StatusBadRequest)
		return
	}

	getPostDynamicsRequest := statspb.GetPostDynamicsRequest{
		PostId: int32(postID),
		Metric: metric,
	}
	getPostDynamicsResp, err := h.statisticsServiceClient.GetPostDynamics(r.Context(), &getPostDynamicsRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get post dynamics: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getPostDynamicsResp)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get post dynamics: error marshalling response: %v", err), http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(
			w, fmt.Sprintf("get post dynamics: error writing response: %v", err), http.StatusInternalServerError,
		)
	}
}
