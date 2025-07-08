package converters

import (
	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/dto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DynamicsToPB(dynamics dto.PostDynamics) []*statspb.PostStat {
	postDynamics := make([]*statspb.PostStat, len(dynamics))
	for i := 0; i < len(dynamics); i++ {
		postDynamics[i] = &statspb.PostStat{
			Date:        timestamppb.New(dynamics[i].Date),
			MetricValue: dynamics[i].MetricValue,
		}
	}

	return postDynamics
}

func TopPostsToPB(top dto.Top) []*statspb.TopPostStat {
	posts := make([]*statspb.TopPostStat, len(top))
	for i := 0; i < len(top); i++ {
		posts[i] = &statspb.TopPostStat{
			PostId:      top[i].ID,
			MetricValue: top[i].MetricValue,
		}
	}

	return posts
}

func TopUsersToPB(top dto.Top) []*statspb.TopUserStat {
	posts := make([]*statspb.TopUserStat, len(top))
	for i := 0; i < len(top); i++ {
		posts[i] = &statspb.TopUserStat{
			UserId:      top[i].ID,
			MetricValue: top[i].MetricValue,
		}
	}

	return posts
}
