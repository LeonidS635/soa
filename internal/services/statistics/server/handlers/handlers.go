package handlers

import (
	"github.com/LeonidS635/soa/internal/pkg/services/statspb"
	"github.com/LeonidS635/soa/internal/services/statistics/usecase"
)

type StatsHandlers struct {
	useCase *usecase.StatsUseCase
	statspb.UnimplementedStatsServiceServer
}

func NewStatsHandlers(useCase *usecase.StatsUseCase) *StatsHandlers {
	return &StatsHandlers{useCase: useCase}
}
