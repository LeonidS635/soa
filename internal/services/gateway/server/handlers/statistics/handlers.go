package stats

import (
	"fmt"

	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GateWayStatisticsHandlers struct {
	statisticsServiceClient statspb.StatsServiceClient
}

func NewGateWayStatisticsHandlers(postsServiceHost string) (*GateWayStatisticsHandlers, error) {
	conn, err := grpc.NewClient(postsServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create posts grpc client: %v", err)
	}

	return &GateWayStatisticsHandlers{
		statisticsServiceClient: statspb.NewStatsServiceClient(conn),
	}, nil
}
