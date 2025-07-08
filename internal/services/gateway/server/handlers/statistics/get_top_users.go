package stats

import (
	"fmt"
	"net/http"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *GateWayStatisticsHandlers) GetTopUsers(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, fmt.Sprintf("get top users: unknown metric: %s", metricStr), http.StatusBadRequest)
		return
	}

	getTopUsersRequest := statspb.GetTopUsersRequest{
		Metric: metric,
	}
	getTopUsersResp, err := h.statisticsServiceClient.GetTopUsers(r.Context(), &getTopUsersRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("get top users: %v", err), http.StatusInternalServerError)
		return
	}

	data, err := protojson.Marshal(getTopUsersResp)
	if err != nil {
		http.Error(
			w, fmt.Sprintf("get top users: error marshalling response: %v", err), http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	written, err := w.Write(data)
	if err != nil || written != len(data) {
		http.Error(
			w, fmt.Sprintf("get top users: error writing response: %v", err), http.StatusInternalServerError,
		)
	}
}
