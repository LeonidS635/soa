package stats

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayStatisticsHandlers) GetTopPosts(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, fmt.Sprintf("get top posts: unknown metric: %s", metricStr), http.StatusBadRequest)
		return
	}

	getTopPostsRequest := statspb.GetTopPostsRequest{
		Metric: metric,
	}
	getTopPostsResp, err := h.statisticsServiceClient.GetTopPosts(r.Context(), &getTopPostsRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get top posts: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getTopPostsResp)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get top posts: error marshalling response: %v", err), http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(
			w, fmt.Sprintf("get top posts: error writing response: %v", err), http.StatusInternalServerError,
		)
	}
}
